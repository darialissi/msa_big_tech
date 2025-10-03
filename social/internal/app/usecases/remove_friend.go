package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) RemoveFriend(fr *dto.RemoveFriend) (*models.UserFriend, error) {

	friend, err := sc.RepoFriend.Delete(fr)

	if err != nil {
		return nil, fmt.Errorf("RemoveFriend: Delete: %w", err)
	}

	if friend == nil {
		return nil, ErrNoFriend
	}

	return friend, nil
}