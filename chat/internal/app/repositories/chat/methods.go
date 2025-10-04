package chat

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (r *Repository) Save(chat *dto.CreateDirectChat) (*models.DirectChat, error) {

	return &models.DirectChat{}, nil
}

func (r *Repository) FetchById(chatId dto.ChatID) (*models.DirectChat, error) {

	return &models.DirectChat{}, nil
}

func (r *Repository) FetchChatMembers(chatId dto.ChatID) ([]*models.ChatParticipant, error) {

	return []*models.ChatParticipant{}, nil
}

func (r *Repository) FetchManyByUserId(userId dto.UserID) ([]*models.DirectChat, error) {

	return []*models.DirectChat{}, nil
}