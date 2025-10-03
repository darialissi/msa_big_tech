package chat

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (r *Repository) Save(chat *dto.CreateChat) (*models.Chat, error) {

	return &models.Chat{}, nil
}

func (r *Repository) FetchById(chatId dto.ChatID) (*models.Chat, error) {

	return &models.Chat{}, nil
}

func (r *Repository) FetchChatMembers(chatId dto.ChatID) ([]*models.ChatParticipant, error) {

	return []*models.ChatParticipant{}, nil
}

func (r *Repository) FetchManyByUserId(userId dto.UserID) ([]*models.Chat, error) {

	return []*models.Chat{}, nil
}