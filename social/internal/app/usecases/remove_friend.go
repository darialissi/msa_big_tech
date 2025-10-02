package usecases

import (
	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) ListFriends(userId dto.UserID) ([]models.FriendID, error) {

	return []models.FriendID{}, nil
}