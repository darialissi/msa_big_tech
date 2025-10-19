package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)

func (sc *SocialUsecase) DeclineFriendRequest(ctx context.Context, reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	// Проверка наличия FriendRequest
	frReq, err := sc.RepoFriendReq.FetchById(ctx, models.FriendRequestID(reqId))

	if err != nil {
		return nil, fmt.Errorf("DeclineFriendRequest: RepoFriendReq.FetchById: %w", err)
	}

	if frReq == nil {
		return nil, ErrNoFriendRequest
	}

	// Изменение статуса на Declined возможно только из Pending
	if frReq.Status != models.FriendRequestStatusPending {
		return nil, ErrBadStatus
	}

	transition := &models.FriendRequest{
		ID:     frReq.ID,
		Status: models.FriendRequestStatusDeclined,
	}

	updated, err := sc.RepoFriendReq.UpdateStatus(ctx, transition)

	if err != nil {
		return nil, fmt.Errorf("DeclineFriendRequest: RepoFriendReq.UpdateStatus: %w", err)
	}

	return updated, nil
}
