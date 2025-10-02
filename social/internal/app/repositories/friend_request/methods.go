package friend_request


import (
	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (r *Repository) Save(req *dto.SendFriendRequest) (models.FriendRequestID, error) {

	return models.FriendRequestID(0), nil
}

func (r *Repository) UpdateStatus(reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	return &models.FriendRequest{}, nil
}

func (r *Repository) FetchById(reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	return &models.FriendRequest{}, nil
}

func (r *Repository) FetchManyByUserId(userId dto.UserID) ([]*models.FriendRequest, error) {

	return []*models.FriendRequest{}, nil
}