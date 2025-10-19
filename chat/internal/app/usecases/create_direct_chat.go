package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)

func (ch *ChatUsecase) CreateDirectChat(ctx context.Context, req *dto.CreateDirectChat) (*models.DirectChat, error) {

	// TODO: проверить существование пользователей

	model := &models.DirectChat{
		CreatorID: models.UserID(req.CreatorID),
	}

	var saved *models.DirectChat

	// Операции сохранения данных в транзакции
	err := ch.txMan.RunReadCommitted(ctx, func(txCtx context.Context) error {
		// Сохранение чата
		res, err := ch.repoChat.Save(ctx, model)

		if err != nil {
			return fmt.Errorf("CreateDirectChat: repoChat.Save: %w", err)
		}

		saved = res

		// Сохранение участников чата
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
			return fmt.Errorf("CreateDirectChat: repoChatMember.SaveMultiple: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return saved, nil
}
