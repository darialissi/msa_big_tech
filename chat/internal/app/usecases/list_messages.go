package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)

func (ch *ChatUsecase) ListMessages(ctx context.Context, lm *dto.ListMessages) ([]*models.Message, *models.Cursor, error) {

	if lm.Limit <= 0 {
		lm.Limit = 10 // дефолтное значение
	}
	if lm.Limit > 10 {
		lm.Limit = 50 // максимальное значение
	}

	cursor := &models.Cursor{
		Limit:      lm.Limit,
		NextCursor: models.MessageID(lm.Cursor),
	}

	chatId := models.ChatID(lm.ChatID)

	messages, nextCursor, err := ch.repoMessage.FetchManyByChatIdCursor(ctx, chatId, cursor)

	if err != nil {
		return nil, nextCursor, fmt.Errorf("ListMessages: repoMessage.FetchManyByChatIdCursor: %w", err)
	}

	if len(messages) == 0 {
		return nil, nextCursor, ErrNoMessages
	}

	return messages, nextCursor, nil
}
