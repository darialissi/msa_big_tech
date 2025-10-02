package social_grpc

import (
	"context"

	"github.com/darialissi/msa_big_tech/social/pkg"
)

func (s *service) SendFriendRequest(ctx context.Context, req *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {

	return &social.SendFriendRequestResponse{FriendRequest: &social.FriendRequest{}}, nil
}

func (s *service) ListRequests(ctx context.Context, req *social.ListFriendRequestsRequest) (*social.ListFriendRequestsResponse, error) {

	return &social.ListFriendRequestsResponse{FriendRequests: []*social.FriendRequest{}}, nil
}

func (s *service) AcceptFriendRequest(ctx context.Context, req *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {

	return &social.AcceptFriendRequestResponse{FriendRequest: &social.FriendRequest{}}, nil
}

func (s *service) DeclineFriendRequest(ctx context.Context, req *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {

	return &social.DeclineFriendRequestResponse{FriendRequest: &social.FriendRequest{}}, nil
}

func (s *service) RemoveFriend(ctx context.Context, req *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {

	return &social.RemoveFriendResponse{}, nil
}

func (s *service) ListFriends(ctx context.Context, req *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {

	return &social.ListFriendsResponse{}, nil
}
