package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) SendMessage(m *dto.SendMessage) (*models.Message, error) {

	res, err := ch.repoMessage.Save(m)

	if err != nil {
		return nil, fmt.Errorf("SendMessage: Save: %w", err)
	}

	return res, nil
}