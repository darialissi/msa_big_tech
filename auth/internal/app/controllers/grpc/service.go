package auth_grpc

import (
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases"
	auth "github.com/darialissi/msa_big_tech/auth/pkg"
)


type service struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	auth.UnimplementedAuthServiceServer

	AuthUsecase *usecases.AuthUsecase
}

func NewServer(authUC *usecases.AuthUsecase) *service {
    return &service{
        AuthUsecase: authUC,
    }
}