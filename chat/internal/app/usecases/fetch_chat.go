package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)

func (ch *ChatUsecase) FetchChat(ctx context.Context, chatId dto.ChatID) (*models.DirectChat, error) {

	res, err := ch.repoChat.FetchById(ctx, models.ChatID(chatId))

	if err != nil {
		return nil, fmt.Errorf("FetchChat: repoChat.FetchById: %w", err)
	}

	if res == nil {
		return nil, ErrNotExistedChat
	}

	return res, nil
}
