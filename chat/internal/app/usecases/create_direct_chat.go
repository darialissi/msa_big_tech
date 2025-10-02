package usecases

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) CreateDirectChat(userId dto.UserID) (models.ChatID, error) {

	return models.ChatID(0), nil
}