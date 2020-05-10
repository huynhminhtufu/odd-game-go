package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/oddx-team/odd-game-server/config"
	pb "github.com/oddx-team/odd-game-server/pb"
	"github.com/oddx-team/odd-game-server/pkg/utils"
	"google.golang.org/grpc"
)

type Server struct {
	http.Server
	cfg          *config.Config
	Addrs        []string // addresses on which the server listens for new connection.
	inShutdown   uint32   // indicates whether the server is in shutdown or not
	requestCount int32    // counter holds no. of request in progress.
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) runGRPCGateway() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = pb.RegisterOddHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", s.cfg.GRPCAddress), opts)
	if err != nil {
		return err
	}

	muxHttp := http.NewServeMux()
	muxHttp.Handle("/", utils.ForwardAccessToken(mux))

	log.Println("Started HTTP server at port " + strconv.Itoa(s.cfg.HttpAddress))
	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.HttpAddress), muxHttp)
}
