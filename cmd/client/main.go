package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/l"
	"github.com/oddx-team/odd-game-server/pb"
	"google.golang.org/grpc"
)

var (
	ll = l.New()
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, fmt.Sprintf(":%v", cfg.GRPCAddress),
		grpc.WithInsecure(), // grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, ""))
	)

	if err != nil {
		ll.Fatal("Failed to dial server:", l.Error(err))
	}

	defer conn.Close()

	c := pb.NewOddClient(conn)
	res, err := c.Liveness(context.Background(), &pb.LivenessRequest{})

	log.Println(res, err)
}
