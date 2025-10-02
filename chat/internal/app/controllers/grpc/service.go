package chat_grpc

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases"
	chat "github.com/darialissi/msa_big_tech/chat/pkg"
)


type service struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	chat.UnimplementedChatServiceServer

	ChatUsecase *usecases.ChatUsecase
}

func NewServer(chatUC *usecases.ChatUsecase) *service {
    return &service{
        ChatUsecase: chatUC,
    }
}