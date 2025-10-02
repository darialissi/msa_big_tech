package main

import (
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chat_grpc "github.com/darialissi/msa_big_tech/chat/internal/app/controllers/grpc"
	chat_repo "github.com/darialissi/msa_big_tech/chat/internal/app/repositories/chat"
	message_repo "github.com/darialissi/msa_big_tech/chat/internal/app/repositories/message"
	chat "github.com/darialissi/msa_big_tech/chat/pkg"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases"
)


func main() {
    // DI
    chatRepo := chat_repo.NewRepository(&sql.DB{})
    messageRepo := message_repo.NewRepository(&sql.DB{})
    
    deps := usecases.Deps{
        RepoChat:  chatRepo,
        RepoMessage: messageRepo,
    }

	if err := deps.Valid(); err != nil {
        log.Fatalf("failed to create chat usecase: %v", err)
	}
    
    chatUC := usecases.NewChatUsecase(deps)

	implementation := chat_grpc.NewServer(chatUC) // наша реализация сервера

	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	chat.RegisterChatServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}