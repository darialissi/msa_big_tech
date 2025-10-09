package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chat_grpc "github.com/darialissi/msa_big_tech/chat/internal/app/controllers/grpc"
	chat_repo "github.com/darialissi/msa_big_tech/chat/internal/app/repositories/chat"
	chat_member_repo "github.com/darialissi/msa_big_tech/chat/internal/app/repositories/chat_member"
	message_repo "github.com/darialissi/msa_big_tech/chat/internal/app/repositories/message"
	chat "github.com/darialissi/msa_big_tech/chat/pkg"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases"
)


func main() {
    // DI
	ctx := context.Background()

	pool, err := NewPostgresConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

    chatRepo := chat_repo.NewRepository(pool)
    chatMemberRepo := chat_member_repo.NewRepository(pool)
    messageRepo := message_repo.NewRepository(pool)
    
    deps := usecases.Deps{
        RepoChat:  chatRepo,
        RepoMessage: messageRepo,
		RepoChatMember: chatMemberRepo,
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