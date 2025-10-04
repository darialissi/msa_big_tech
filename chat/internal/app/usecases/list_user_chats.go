package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) ListUserChats(userId dto.UserID) ([]*models.DirectChat, error) {

	res, err := ch.repoChat.FetchManyByUserId(userId)

	if err != nil {
		return nil, fmt.Errorf("ListUserChats: FetchManyByUserId : %w", err)
	}

	if len(res) == 0 {
		return nil, ErrNoUserChats
	}

	return res, nil
}