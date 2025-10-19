package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)

func (sc *SocialUsecase) ListFriendRequests(ctx context.Context, userId dto.UserID) ([]*models.FriendRequest, error) {

	frReqs, err := sc.RepoFriendReq.FetchManyByUserId(ctx, models.UserID(userId))

	if err != nil {
		return nil, fmt.Errorf("ListFriendRequests: RepoFriendReq.FetchManyByUserId: %w", err)
	}

	if len(frReqs) == 0 {
		return nil, ErrUserNoFriendRequestsIn
	}

	return frReqs, nil
}
