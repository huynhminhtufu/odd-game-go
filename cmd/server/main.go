package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/oddx-team/odd-game-server/config"
	pb "github.com/oddx-team/odd-game-server/pb"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	//grpc server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)

	// Register Prometheus metrics handler.
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)

	svc := registerService(cfg)
	pb.RegisterOddServer(s, svc)

	// handle signal
	_, ctxCancel := context.WithCancel(context.Background())
	go func() {
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
		<-osSignal
		ctxCancel()
		// Wait for maximum 15s
		go func() {
			var durationSec time.Duration = 15
			timer := time.NewTimer(durationSec * time.Second)
			<-timer.C
			log.Fatal("Force shutdown due to timeout!")
		}()
	}()

	go func() {
		gw := NewServer(cfg)
		fmt.Println("HTTP server start listening", fmt.Sprintf("HTTP address: %d", cfg.HttpAddress))
		err := gw.runGRPCGateway()
		if err != nil {
			log.Println("Cannot start server")
			return
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCAddress))
	log.Println("Started gRPC server at port " + strconv.Itoa(cfg.GRPCAddress))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
