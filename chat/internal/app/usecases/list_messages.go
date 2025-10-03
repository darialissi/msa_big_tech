package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) ListMessages(chatId dto.ChatID) ([]*models.Message, error) {

	res, err := ch.repoMessage.FetchManyByChatId(chatId)

	if err != nil {
		return nil, fmt.Errorf("ListMessages: FetchManyByChatId: %w", err)
	}

	if len(res) == 0 {
		return nil, ErrNoMessages
	}

	return res, nil
}