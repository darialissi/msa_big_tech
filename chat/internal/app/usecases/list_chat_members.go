package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) ListChatMembers(chatId dto.ChatID) ([]*models.ChatParticipant, error) {

	res, err := ch.repoChat.FetchChatMembers(chatId)

	if err != nil {
		return nil, fmt.Errorf("ListChatMembers: FetchChatMembers: %w", err)
	}

	return res, nil
}