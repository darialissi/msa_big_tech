package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) ListFriends(ctx context.Context, userId dto.UserID) ([]*models.UserFriend, error) {

	friends, err := sc.RepoFriend.FetchManyByUserId(ctx, models.UserID(userId))

	if err != nil {
		return nil, fmt.Errorf("ListFriends: RepoFriend.FetchManyByUserId: %w", err)
	}

	if len(friends) == 0 {
		return nil, ErrUserNoFriends
	}

	return friends, nil
}