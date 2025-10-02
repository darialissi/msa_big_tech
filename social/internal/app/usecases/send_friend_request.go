package usecases

import (
	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) SendFriendRequest(req *dto.SendFriendRequest) (models.FriendRequestID, error) {

	return models.FriendRequestID(0), nil
}