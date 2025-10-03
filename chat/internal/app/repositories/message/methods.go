package message

import (
	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (r *Repository) Save(message *dto.SendMessage) (*models.Message, error) {

	return &models.Message{}, nil
}

func (r *Repository) FetchById(messageId dto.MessageID) (*models.Message, error) {

	return &models.Message{}, nil
}

func (r *Repository) FetchManyByChatId(chatId dto.ChatID) ([]*models.Message, error) {

	return []*models.Message{}, nil
}