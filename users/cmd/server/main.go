package main

import (
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	user_repo "github.com/darialissi/msa_big_tech/users/internal/app/repositories/user"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases"
	users_grpc "github.com/darialissi/msa_big_tech/users/internal/app/controllers/grpc"
	users "github.com/darialissi/msa_big_tech/users/pkg"
)


func main() {
    // DI
	ctx := context.Background()

	pool, err := NewPostgresConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

    usersRepo := user_repo.NewRepository(pool)
    
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