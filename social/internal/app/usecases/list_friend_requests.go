package usecases

import (
	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) ListFriendRequests(userId dto.UserID) ([]*models.FriendRequest, error) {

	return []*models.FriendRequest{}, nil
}