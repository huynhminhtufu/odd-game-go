package main

import (
	"context"
	"fmt"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/internal/server"
	"github.com/oddx-team/odd-game-server/internal/services"
	"github.com/oddx-team/odd-game-server/pb"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	// gRPC server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpcprometheus.UnaryServerInterceptor,
		)),
	)

	// Register Prometheus metrics handler
	grpcprometheus.EnableHandlingTimeHistogram()
	grpcprometheus.Register(s)

	svc := services.NewService(cfg)
	pb.RegisterOddServer(s, svc)

	// Handle signal
	_, ctxCancel := context.WithCancel(context.Background())
	go func() {
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
		<-osSignal
		ctxCancel()
		// Wait for maximum 15s
		go func() {
			var duration time.Duration = 15
			timer := time.NewTimer(duration * time.Second)
			<-timer.C
			log.Fatal("Force shutdown due to timeout")
		}()
	}()

	go func() {
		gw := server.NewServer(cfg)
		log.Println("Started HTTP server at port " + strconv.Itoa(cfg.HttpAddress))
		err := gw.RunGRPCGateway()
		if err != nil {
			log.Println("Cannot start server")
			return
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCAddress))
	log.Println("Started gRPC server at port " + strconv.Itoa(cfg.GRPCAddress))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
