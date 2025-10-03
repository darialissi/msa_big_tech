package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) SendFriendRequest(req *dto.SendFriendRequest) (*models.FriendRequest, error) {

	reqFr := &dto.SaveFriendRequest{
		StatusID: dto.FriendRequestStatus(models.FriendRequestStatusPending),
		FromUserID: req.FromUserID,
		ToUserID: req.ToUserID,
	}

	frReq, err := sc.RepoFriendReq.Save(reqFr)

	if err != nil {
		return nil, fmt.Errorf("SendFriendRequest: Save: %w", err)
	}

	return frReq, nil
}