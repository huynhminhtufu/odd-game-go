package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/internal/chat"
	"github.com/oddx-team/odd-game-server/internal/user"
	pb "github.com/oddx-team/odd-game-server/pb"
	"github.com/oddx-team/odd-game-server/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
)

func runGRPCGateway(grpcPort, httpPort int) (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = pb.RegisterChatHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		return err
	}
	err = pb.RegisterUserHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		return err
	}

	muxHttp := http.NewServeMux()
	muxHttp.Handle("/", utils.ForwardAccessToken(mux))

	log.Println("Started HTTP server at port " + strconv.Itoa(httpPort))
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), muxHttp)
}

func main() {
	cfg := config.Load()

	go func() {
		err := runGRPCGateway(cfg.GrpcPort, cfg.HttpPort)
		if err != nil {
			log.Println("Cannot start server")
			return
		}
	}()
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GrpcPort))
	log.Println("Started gRPC server at port " + strconv.Itoa(cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterChatServer(s, &chat.Server{})
	pb.RegisterUserServer(s, &user.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}