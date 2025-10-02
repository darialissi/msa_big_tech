package usecases

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) SendMessage(m *dto.SendMessage) (*models.Message, error) {

	return &models.Message{}, nil
}