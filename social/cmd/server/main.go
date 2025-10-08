package main

import (
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/darialissi/msa_big_tech/social/internal/app/usecases"
	friend_repo "github.com/darialissi/msa_big_tech/social/internal/app/repositories/friend"
	friend_req_repo "github.com/darialissi/msa_big_tech/social/internal/app/repositories/friend_request"
	social_grpc "github.com/darialissi/msa_big_tech/social/internal/app/controllers/grpc"
	social "github.com/darialissi/msa_big_tech/social/pkg"
)


func main() {
    // DI
	ctx := context.Background()

	pool, err := NewPostgresConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

    friendRepo := friend_repo.NewRepository(pool)
    friendReqRepo := friend_req_repo.NewRepository(pool)
    
    deps := usecases.Deps{
        RepoFriend:  friendRepo,
        RepoFriendReq: friendReqRepo,
    }
    
    socialUC, err := usecases.NewSocialUsecase(deps)
    if err != nil {
        log.Fatalf("failed to create social usecase: %v", err)
    }
	
    implementation := social_grpc.NewServer(socialUC) // наша реализация сервера

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	social.RegisterSocialServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}