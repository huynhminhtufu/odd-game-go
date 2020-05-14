package services

import (
	"context"

	"github.com/oddx-team/odd-game-server/pb"
)

func (s *service) GetChats(_ context.Context, _ *pb.GetChatsRequest) (*pb.GetChatsResponse, error) {
	return &pb.GetChatsResponse{
		Chats: nil,
	}, nil
}
