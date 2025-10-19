package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)

func (sc *SocialUsecase) SendFriendRequest(ctx context.Context, req *dto.SendFriendRequest) (*models.FriendRequest, error) {

	if req.FromUserID == req.ToUserID {
		return nil, ErrRequestToYourself
	}

	// TODO: проверить сущестование пользователя через клиента Users Service

	reqFr := &models.FriendRequest{
		Status:     models.FriendRequestStatusPending,
		FromUserID: models.UserID(req.FromUserID),
		ToUserID:   models.UserID(req.ToUserID),
	}

	frReq, err := sc.RepoFriendReq.Save(ctx, reqFr)

	if err != nil {
		return nil, fmt.Errorf("SendFriendRequest: RepoFriendReq.Save: %w", err)
	}

	return frReq, nil
}
