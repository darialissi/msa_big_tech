package main

import (
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	users_repo "github.com/darialissi/msa_big_tech/users/internal/app/repositories/users"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases"
	users_grpc "github.com/darialissi/msa_big_tech/users/internal/app/controllers/grpc"
	users "github.com/darialissi/msa_big_tech/users/pkg"
)


func main() {
    // DI
    usersRepo := users_repo.NewRepository(&sql.DB{})
    
    usersUC := usecases.NewUsersUsecase(usersRepo)
	
    implementation := users_grpc.NewServer(usersUC) // наша реализация сервера

	lis, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	users.RegisterUsersServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}