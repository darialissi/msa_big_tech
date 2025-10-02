package usecases

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) ListChatMembers(chatId dto.ChatID) ([]models.UserID, error) {

	return []models.UserID{}, nil
}