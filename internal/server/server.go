package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/pb"
	"github.com/oddx-team/odd-game-server/pkg/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net/http"
)

type Server struct {
	http.Server
	cfg          *config.Config
	Addrs        []string // Addresses on which the server listens for new connection
	inShutdown   uint32   // Indicates whether the server is in shutdown or not
	requestCount int32    // Counter holds no. of request in progress
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) RunGRPCGateway() (err error) {
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
	muxHttp.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})
	muxHttp.Handle("/", utils.ForwardAccessToken(mux))

	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.HttpAddress), muxHttp)
}