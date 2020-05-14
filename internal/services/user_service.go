package services

import (
	"context"

	"github.com/oddx-team/odd-game-server/pb"
)

func (s *service) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	email := req.Email
	return &pb.GetUserResponse{
		User: &pb.UserEntity{
			Email: email,
			Name:  email,
		},
	}, nil
}

func (s *service) GetUsers(_ context.Context, _ *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{
		Users: nil,
	}, nil
}
