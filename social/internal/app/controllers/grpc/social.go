package social_grpc

import (
	"context"

	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
	social "github.com/darialissi/msa_big_tech/social/pkg"
)

func (s *service) SendFriendRequest(ctx context.Context, req *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {

    // TODO: получить id из jwt MW
	userId := "00000000-0000-0000-0000-0000000000000"

	form := &dto.SendFriendRequest{
		FromUserID: dto.UserID(userId),
		ToUserID: dto.UserID(req.UserId),
	}

	res, err := s.SocialUsecase.SendFriendRequest(form)

	if err != nil {
		return nil, err
	}

	resp := &social.FriendRequest{
		RequestId: string(res.ID),
		Status: social.Status(res.Status),
		FromUserId: string(res.FromUserID),
		ToUserId: string(res.ToUserID),
	}

	return &social.SendFriendRequestResponse{FriendRequest: resp}, nil
}

func (s *service) ListRequests(ctx context.Context, req *social.ListFriendRequestsRequest) (*social.ListFriendRequestsResponse, error) {

	res, err := s.SocialUsecase.ListFriendRequests(dto.UserID(req.UserId))

	if err != nil {
		return nil, err
	}

	resp := make([]*social.FriendRequest, len(res))

    for i, r := range res {
        resp[i] = &social.FriendRequest{
			RequestId: string(r.ID),
			Status: social.Status(r.Status),
			FromUserId: string(r.FromUserID),
			ToUserId: string(r.ToUserID),
		}
	}

	return &social.ListFriendRequestsResponse{FriendRequests: resp}, nil
}

func (s *service) AcceptFriendRequest(ctx context.Context, req *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {

	res, err := s.SocialUsecase.AcceptFriendRequest(dto.FriendRequestID(req.FriendRequestId))

	if err != nil {
		return nil, err
	}

	resp := &social.FriendRequest{
		RequestId: string(res.ID),
		Status: social.Status(res.Status),
		FromUserId: string(res.FromUserID),
		ToUserId: string(res.ToUserID),
	}

	return &social.AcceptFriendRequestResponse{FriendRequest: resp}, nil
}

func (s *service) DeclineFriendRequest(ctx context.Context, req *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {

	res, err := s.SocialUsecase.DeclineFriendRequest(dto.FriendRequestID(req.FriendRequestId))

	if err != nil {
		return nil, err
	}

	resp := &social.FriendRequest{
		RequestId: string(res.ID),
		Status: social.Status(res.Status),
		FromUserId: string(res.FromUserID),
		ToUserId: string(res.ToUserID),
	}

	return &social.DeclineFriendRequestResponse{FriendRequest: resp}, nil
}

func (s *service) RemoveFriend(ctx context.Context, req *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {

    // TODO: получить id из jwt MW
	userId := "00000000-0000-0000-0000-0000000000000"

	form := &dto.RemoveFriend{
		UserID: dto.UserID(userId),
		FriendID: dto.FriendID(req.FriendId),
	}

	res, err := s.SocialUsecase.RemoveFriend(form)

	if err != nil {
		return nil, err
	}

	return &social.RemoveFriendResponse{FriendId: string(res.FriendId)}, nil
}

func (s *service) ListFriends(ctx context.Context, req *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {

	// TODO: реализовать пагинацию

    // TODO: получить id из jwt MW
	userId := "00000000-0000-0000-0000-0000000000000"

	res, err := s.SocialUsecase.ListFriends(dto.UserID(userId))

	if err != nil {
		return nil, err
	}

    resp := make([]string, len(res))

    for i, v := range res {
        resp[i] = string(v.FriendId)
    }

	return &social.ListFriendsResponse{FriendIds: resp}, nil
}
