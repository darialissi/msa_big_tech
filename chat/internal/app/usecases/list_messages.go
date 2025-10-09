package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) ListMessages(ctx context.Context, chatId dto.ChatID) ([]*models.Message, error) {

	res, err := ch.repoMessage.FetchManyByChatId(ctx, models.ChatID(chatId))

	if err != nil {
		return nil, fmt.Errorf("ListMessages: repoMessage.FetchManyByChatId: %w", err)
	}

	if len(res) == 0 {
		return nil, ErrNoMessages
	}

	return res, nil
}