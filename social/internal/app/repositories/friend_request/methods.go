package friend_request


import (
	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (r *Repository) Save(req *dto.SaveFriendRequest) (*models.FriendRequest, error) {

	return &models.FriendRequest{}, nil
}

func (r *Repository) UpdateStatus(req *dto.ChangeStatus) (*models.FriendRequest, error) {

	return &models.FriendRequest{}, nil
}

func (r *Repository) FetchById(reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	return &models.FriendRequest{}, nil
}

func (r *Repository) FetchManyByUserId(userId dto.UserID) ([]*models.FriendRequest, error) {

	return []*models.FriendRequest{}, nil
}