package usecases

import (
	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) DeclineFriendRequest(reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	return &models.FriendRequest{}, nil
}