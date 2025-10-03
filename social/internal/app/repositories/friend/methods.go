package friend

import (
	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (r *Repository) Save(fr *dto.SaveFriend) (*models.UserFriend, error) {

	return &models.UserFriend{}, nil
}

func (r *Repository) Delete(fr *dto.RemoveFriend) (*models.UserFriend, error) {

	return &models.UserFriend{}, nil
}

func (r *Repository) FetchManyByUserId(userId dto.UserID) ([]*models.UserFriend, error) {

	return []*models.UserFriend{}, nil
}