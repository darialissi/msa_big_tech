package usecases

import (
	"context"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) StreamMessages(ctx context.Context, chatId dto.ChatID) ([]*models.Message, error) {

	return []*models.Message{}, nil
}