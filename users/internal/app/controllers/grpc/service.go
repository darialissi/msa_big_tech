package users_grpc

import (
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases"
	users "github.com/darialissi/msa_big_tech/users/pkg"
)

type service struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	users.UnimplementedUsersServiceServer

	UsersUsecase usecases.UsersUsecases
}

func NewServer(usersUC *usecases.UsersUsecase) *service {
	return &service{
		UsersUsecase: usersUC,
	}
}
