package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) DeclineFriendRequest(ctx context.Context, reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	frReq, err := sc.RepoFriendReq.FetchById(ctx, reqId)

	if err != nil {
		return nil, fmt.Errorf("DeclineFriendRequest: RepoFriendReq.FetchById: %w", err)
	}

	if frReq == nil {
		return nil, ErrNoFriendRequest
	}

	// изменение статуса на Declined возможно только из Pending
	if frReq.Status != models.FriendRequestStatusPending {
		return nil, ErrBadStatus
	}

	transition := &dto.UpdateFriendRequest{
		ReqID: reqId,
		Status: dto.FriendRequestStatus(models.FriendRequestStatusDeclined),
	}

	res, err := sc.RepoFriendReq.UpdateStatus(ctx, transition)

	if err != nil {
		return nil, fmt.Errorf("DeclineFriendRequest: RepoFriendReq.UpdateStatus: %w", err)
	}

	return res, nil
}