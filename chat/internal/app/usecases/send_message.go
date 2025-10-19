package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)

func (ch *ChatUsecase) SendMessage(ctx context.Context, d *dto.SendMessage) (*models.Message, error) {

	model := &models.Message{
		ChatID:   models.ChatID(d.ChatID),
		SenderID: models.UserID(d.SenderID),
		Text:     d.Text,
	}

	res, err := ch.repoMessage.Save(ctx, model)

	if err != nil {
		return nil, fmt.Errorf("SendMessage: Save: %w", err)
	}

	return res, nil
}
