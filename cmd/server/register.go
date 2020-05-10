package main

import (
	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/internal/services"
	"github.com/oddx-team/odd-game-server/pb"
)

func registerService(cfg *config.Config) pb.OddServer {
	return services.New(cfg)
}
