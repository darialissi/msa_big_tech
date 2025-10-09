package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) ListChatMembers(ctx context.Context, chatId dto.ChatID) ([]*models.ChatMember, error) {

	res, err := ch.repoChatMember.FetchManyByChatId(ctx, models.ChatID(chatId))

	if err != nil {
		return nil, fmt.Errorf("ListChatMembers: repoChatMember.FetchManyByChatId: %w", err)
	}

	return res, nil
}