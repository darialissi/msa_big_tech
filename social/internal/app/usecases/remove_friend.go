package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) RemoveFriend(ctx context.Context, fr *dto.RemoveFriend) (*models.UserFriend, error) {

	friend, err := sc.RepoFriend.Delete(ctx, fr)

	if err != nil {
		return nil, fmt.Errorf("RemoveFriend: RepoFriend.Delete: %w", err)
	}

	return friend, nil
}