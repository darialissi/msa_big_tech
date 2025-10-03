package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) DeclineFriendRequest(reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	frReq, err := sc.RepoFriendReq.FetchById(reqId)

	if err != nil {
		return nil, fmt.Errorf("DeclineFriendRequest: FetchById: %w", err)
	}

	if frReq == nil {
		return nil, ErrNoFriendRequest
	}

	// изменение статуса на Declined возможно только из Pending
	if frReq.Status != models.FriendRequestStatusPending {
		return nil, ErrBadStatus
	}

	transition := &dto.ChangeStatus{
		ReqID: reqId,
		StatusID: dto.FriendRequestStatus(models.FriendRequestStatusDeclined),
	}

	res, err := sc.RepoFriendReq.UpdateStatus(transition)

	if err != nil {
		return nil, fmt.Errorf("DeclineFriendRequest: UpdateStatus: %w", err)
	}

	return res, nil
}