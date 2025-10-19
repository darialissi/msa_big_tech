package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)

func (sc *SocialUsecase) AcceptFriendRequest(ctx context.Context, reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	// Проверка наличия FriendRequest
	frReq, err := sc.RepoFriendReq.FetchById(ctx, models.FriendRequestID(reqId))

	if err != nil {
		return nil, fmt.Errorf("AcceptFriendRequest: RepoFriendReq.FetchById: %w", err)
	}

	if frReq == nil {
		return nil, ErrNoFriendRequest
	}

	// Изменение статуса на Accepted возможно только из Pending
	if frReq.Status != models.FriendRequestStatusPending {
		return nil, ErrBadStatus
	}

	var updated *models.FriendRequest

	// Операции изменения данных в транзакции
	err = sc.TxMan.RunReadCommitted(ctx, func(txCtx context.Context) error {
		transition := &models.FriendRequest{
			ID:     frReq.ID,
			Status: models.FriendRequestStatusAccepted,
		}

		var err error

		// Обновление статуса заявки
		updated, err = sc.RepoFriendReq.UpdateStatus(txCtx, transition)
		if err != nil {
			return fmt.Errorf("AcceptFriendRequest: RepoFriendReq.UpdateStatus: %w", err)
		}

		friend := &models.UserFriend{
			UserID:   frReq.FromUserID,
			FriendID: frReq.ToUserID,
		}

		// Сохранение Друга
		if _, err := sc.RepoFriend.Save(txCtx, friend); err != nil {
			return fmt.Errorf("AcceptFriendRequest: RepoFriend.Save: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updated, nil
}
