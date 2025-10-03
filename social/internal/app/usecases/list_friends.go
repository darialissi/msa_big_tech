package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) ListFriends(userId dto.UserID) ([]*models.UserFriend, error) {

	friends, err := sc.RepoFriend.FetchManyByUserId(userId)

	if err != nil {
		return nil, fmt.Errorf("ListFriends: FetchManyByUserId: %w", err)
	}

	if len(friends) == 0 {
		return nil, ErrUserNoFriends
	}

	return friends, nil
}