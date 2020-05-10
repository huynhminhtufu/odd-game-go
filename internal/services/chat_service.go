package services

import (
	"context"

	pb "github.com/oddx-team/odd-game-server/pb"
)

func (s *service) GetChats(context context.Context, req *pb.GetChatsRequest) (*pb.GetChatsResponse, error) {
	return &pb.GetChatsResponse{
		Chats: nil,
	}, nil
}
