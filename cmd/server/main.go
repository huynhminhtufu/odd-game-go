package main

import (
	"context"
	"fmt"
	"github.com/oddx-team/odd-game-server/internal/websocket"
	"github.com/oddx-team/odd-game-server/pkg/l"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/internal/server"
	"github.com/oddx-team/odd-game-server/internal/services"
	"github.com/oddx-team/odd-game-server/pb"

	"google.golang.org/grpc"
)

var (
	ll = l.New()
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
		// Wait for maximum 1s
		go func() {
			var duration time.Duration = 1
			timer := time.NewTimer(duration * time.Second)
			<-timer.C
			ll.Fatal("Force shutdown due to timeout")
		}()
	}()

	// Websocket module init
	wsHub := websocket.NewHub(svc)
	go wsHub.Run()

	go func() {
		wsHandler := func(w http.ResponseWriter, req *http.Request) {
			websocket.ServeWs(wsHub, w, req)
		}
		gw := server.NewServer(cfg, wsHandler)
		ll.Info("Started HTTP server at port " + strconv.Itoa(cfg.HttpAddress))
		err := gw.RunGRPCGateway()
		if err != nil {
			ll.Error("Cannot start server")
			return
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCAddress))
	ll.Info("Started gRPC server at port " + strconv.Itoa(cfg.GRPCAddress))
	if err != nil {
		ll.Error("Failed to listen: " + err.Error())
	}

	if err := s.Serve(lis); err != nil {
		ll.Error("Failed to serve: " + err.Error())
	}
}
