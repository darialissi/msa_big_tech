package usecases

import (
	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) RemoveFriend(m *dto.RemoveFriend) (models.FriendID, error) {

	return models.FriendID(0), nil
}