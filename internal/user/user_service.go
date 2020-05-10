package user

import (
	"context"
	pb "github.com/oddx-team/odd-game-server/pb"
)

type Server struct {
	pb.UnimplementedUserServer
}

func (s *Server) GetUser(_ context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	email := in.Email
	return &pb.GetUserResponse{
		User: &pb.UserEntity{
			Email: email,
			Name: email,
		},
	}, nil
}

func (s *Server) GetUsers(_ context.Context, _ *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{
		Users: nil,
	}, nil
}