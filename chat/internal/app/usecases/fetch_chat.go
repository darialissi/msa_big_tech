package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) FetchChat(chatId dto.ChatID) (*models.Chat, error) {

	res, err := ch.repoChat.FetchById(chatId)

	if err != nil {
		return nil, fmt.Errorf("FetchChat: FetchById: %w", err)
	}

	if res == nil {
		return nil, ErrNotExistedChat
	}

	return res, nil
}