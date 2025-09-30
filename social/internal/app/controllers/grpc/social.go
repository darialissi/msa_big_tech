package grpc_hd

import (
	"context"

	"github.com/darialissi/msa_big_tech/social/pkg"
)

type server struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	social.UnimplementedSocialServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SendFriendRequest(ctx context.Context, req *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {

	return &social.SendFriendRequestResponse{FriendRequest: &social.FriendRequest{RequestId: 0, Status: social.Status_STATUS_PENDING}}, nil
}

func (s *server) ListRequests(ctx context.Context, req *social.ListRequestsRequest) (*social.ListRequestsResponse, error) {

	return &social.ListRequestsResponse{Requests: []*social.FriendRequest{}}, nil
}

func (s *server) AcceptFriendRequest(ctx context.Context, req *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {

	return &social.AcceptFriendRequestResponse{FriendRequest: &social.FriendRequest{RequestId: req.RequestId, Status: social.Status_STATUS_APPROVED}}, nil
}

func (s *server) DeclineFriendRequest(ctx context.Context, req *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {

	return &social.DeclineFriendRequestResponse{FriendRequest: &social.FriendRequest{RequestId: req.RequestId, Status: social.Status_STATUS_DECLINED}}, nil
}

func (s *server) RemoveFriend(ctx context.Context, req *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {

	return &social.RemoveFriendResponse{}, nil
}

func (s *server) ListFriends(ctx context.Context, req *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {

	return &social.ListFriendsResponse{FriendUserIds: []uint64{0, 10}}, nil
}
