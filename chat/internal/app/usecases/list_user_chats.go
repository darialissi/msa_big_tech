package usecases

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) ListUserChats(userId dto.UserID) ([]*models.Chat, error) {

	return []*models.Chat{}, nil
}