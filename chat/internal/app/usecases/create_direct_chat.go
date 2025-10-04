package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) CreateDirectChat(chat *dto.CreateDirectChat) (*models.DirectChat, error) {

	res, err := ch.repoChat.Save(chat)

	if err != nil {
		return nil, fmt.Errorf("CreateDirectChat: Save: %w", err)
	}

	return res, nil
}