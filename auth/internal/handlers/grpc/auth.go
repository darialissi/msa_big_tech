package grpc_hd

import (
	"context"

	"github.com/darialissi/msa_big_tech/auth/pkg"
)

type server struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	auth.UnimplementedAuthServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {

	return &auth.RegisterResponse{UserId: 0}, nil
}

func (s *server) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {

	return &auth.LoginResponse{UserId: 0, AccessToken: "", RefreshToken: ""}, nil
}

func (s *server) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {

	return &auth.RefreshResponse{UserId: 0, AccessToken: "", RefreshToken: ""}, nil
}
