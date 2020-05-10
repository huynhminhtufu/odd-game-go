package services

import (
	"context"

	pb "github.com/oddx-team/odd-game-server/pb"
)

func (s *service) GetUser(context context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	email := req.Email
	return &pb.GetUserResponse{
		User: &pb.UserEntity{
			Email: email,
			Name:  email,
		},
	}, nil
}

func (s *service) GetUsers(context context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{
		Users: nil,
	}, nil
}
