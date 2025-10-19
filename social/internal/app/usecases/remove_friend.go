package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)

func (sc *SocialUsecase) RemoveFriend(ctx context.Context, fr *dto.RemoveFriend) (*models.UserFriend, error) {

	in := &models.UserFriend{
		UserID:   models.UserID(fr.UserID),
		FriendID: models.UserID(fr.FriendID),
	}

	friend, err := sc.RepoFriend.Delete(ctx, in)

	if err != nil {
		return nil, fmt.Errorf("RemoveFriend: RepoFriend.Delete: %w", err)
	}

	return friend, nil
}
