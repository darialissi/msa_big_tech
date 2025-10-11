package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) StreamMessages(ctx context.Context, st *dto.StreamMessages) ([]*models.Message, error) {

	res, err := ch.repoMessage.StreamMany(ctx, models.ChatID(st.ChatID), time.Unix(st.SinceUnix, 0))

	if err != nil {
		return nil, fmt.Errorf("StreamMessages: repoMessage.StreamMany: %w", err)
	}

	if len(res) == 0 {
		return nil, ErrNoMessages
	}

	return res, nil
}