package chat

import (
	"context"
	pb "github.com/oddx-team/odd-game-server/pb"
)

type Server struct {
	pb.UnimplementedChatServer
}

func (s *Server) GetChats(_ context.Context, _ *pb.GetChatsRequest) (*pb.GetChatsResponse, error) {
	return &pb.GetChatsResponse{
		Chats: nil,
	}, nil
}