package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) ListUserChats(ctx context.Context, userId dto.UserID) ([]*models.ChatMember, error) {

	res, err := ch.repoChatMember.FetchManyByUserId(ctx, models.UserID(userId))

	if err != nil {
		return nil, fmt.Errorf("ListUserChats: repoChatMember.FetchManyByUserId: %w", err)
	}

	if len(res) == 0 {
		return nil, ErrNoUserChats
	}

	return res, nil
}