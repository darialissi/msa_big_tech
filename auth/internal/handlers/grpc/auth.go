package grpc_hd

import (
	"context"

	pb "msa_big_tech/auth/pkg/api/proto"
)

type server struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	pb.UnimplementedAuthServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	return &pb.RegisterResponse{UserId: 0}, nil
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	return &pb.LoginResponse{UserId: 0, AccessToken: "", RefreshToken: ""}, nil
}

func (s *server) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {

	return &pb.RefreshResponse{UserId: 0, AccessToken: "", RefreshToken: ""}, nil
}
