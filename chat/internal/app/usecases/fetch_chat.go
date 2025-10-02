package usecases

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) FetchChat(chatId dto.ChatID) (*models.Chat, error) {

	return &models.Chat{}, nil
}