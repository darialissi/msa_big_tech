package grpc_hd

import (
	"context"

	"msa_big_tech/users/pkg/api/proto"
)

type server struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	users.UnimplementedUsersServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) CreateProfile(ctx context.Context, req *users.CreateProfileRequest) (*users.CreateProfileResponse, error) {
	if err := req.Validate(); err != nil { // не генерится правило валидации в pb :((
		return nil, err
	}

	return &users.CreateProfileResponse{UserProfile: &users.UserProfile{Nickname: req.Nickname}}, nil
}

func (s *server) UpdateProfile(ctx context.Context, req *users.UpdateProfileRequest) (*users.UpdateProfileResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &users.UpdateProfileResponse{UserProfile: &users.UserProfile{Nickname: req.Nickname}}, nil
}

func (s *server) GetProfileByID(ctx context.Context, req *users.GetProfileByIDRequest) (*users.GetProfileByIDResponse, error) {

	return &users.GetProfileByIDResponse{UserProfile: &users.UserProfile{}}, nil
}

func (s *server) GetProfileByNickname(ctx context.Context, req *users.GetProfileByNicknameRequest) (*users.GetProfileByNicknameResponse, error) {

	return &users.GetProfileByNicknameResponse{UserProfile: &users.UserProfile{Nickname: req.Nickname}}, nil
}

func (s *server) SearchByNickname(ctx context.Context, req *users.SearchByNicknameRequest) (*users.SearchByNicknameResponse, error) {

	return &users.SearchByNicknameResponse{Profiles: []*users.UserProfile{&users.UserProfile{}}}, nil
}