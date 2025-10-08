package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) AcceptFriendRequest(ctx context.Context, reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	frReq, err := sc.RepoFriendReq.FetchById(ctx, reqId)

	if err != nil {
		return nil, fmt.Errorf("AcceptFriendRequest: RepoFriendReq.FetchById: %w", err)
	}

	if frReq == nil {
		return nil, ErrNoFriendRequest
	}

	// изменение статуса на Accepted возможно только из Pending
	if frReq.Status != models.FriendRequestStatusPending {
		return nil, ErrBadStatus
	}

	transition := &dto.UpdateFriendRequest{
		ReqID: reqId,
		Status: dto.FriendRequestStatus(models.FriendRequestStatusAccepted),
	}

	updated, err := sc.RepoFriendReq.UpdateStatus(ctx, transition)

	if err != nil {
		return nil, fmt.Errorf("AcceptFriendRequest: RepoFriendReq.UpdateStatus: %w", err)
	}

	// добавление Друга
	friend := &dto.SaveFriend{
		UserID: dto.UserID(frReq.FromUserID),
		FriendID: dto.UserID(frReq.ToUserID),
	}

	if _, err := sc.RepoFriend.Save(ctx, friend); err != nil {
		return nil, fmt.Errorf("AcceptFriendRequest: RepoFriend.Save: %w", err)
	}

	return updated, nil
}