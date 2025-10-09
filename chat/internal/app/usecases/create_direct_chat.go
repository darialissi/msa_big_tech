package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


func (ch *ChatUsecase) CreateDirectChat(ctx context.Context, req *dto.CreateDirectChat) (*models.DirectChat, error) {

	model := &models.DirectChat{
		CreatorID: models.UserID(req.CreatorID),
	}

	// сохраняем чат
	res, err := ch.repoChat.Save(ctx, model)

	if err != nil {
		return nil, fmt.Errorf("CreateDirectChat: repoChat.Save: %w", err)
	}

	// сохраняем участников чата
	members := make([]*models.ChatMember, 2)
	
	members[0] = &models.ChatMember{
		ChatID: res.ID,
		UserID: models.UserID(req.CreatorID),
	}

	members[1] = &models.ChatMember{
		ChatID: res.ID,
		UserID: models.UserID(req.MemberID),
	}

	if _, err = ch.repoChatMember.SaveMultiple(ctx, members); err != nil {
		return nil, fmt.Errorf("CreateDirectChat: repoChatMember.SaveMultiple: %w", err)
	}

	return res, nil
}