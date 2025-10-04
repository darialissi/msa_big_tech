package users_grpc

import (
	"context"

	"buf.build/go/protovalidate"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
	users "github.com/darialissi/msa_big_tech/users/pkg"
)


func (s *service) CreateProfile(ctx context.Context, req *users.CreateProfileRequest) (*users.CreateProfileResponse, error) {

	v, err := protovalidate.New()
    if err != nil {
		return nil, models.ErrValidationFailed
    }

    if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

    // TODO: получить id из jwt MW
	userId := "00000000-0000-0000-0000-0000000000000"

	form := &dto.CreateUser{
		ID: dto.UserID(userId),
		Nickname: req.Nickname,
		Bio: req.Bio,
		Avatar: dto.Url(req.AvatarUrl),
	}

	res, err := s.UsersUsecase.CreateUser(form)

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(res.ID),
		Nickname: res.Nickname,
		Bio: res.Bio,
		AvatarUrl: res.Avatar,
	}

	return &users.CreateProfileResponse{UserProfile: userProfile}, nil
}

func (s *service) UpdateProfile(ctx context.Context, req *users.UpdateProfileRequest) (*users.UpdateProfileResponse, error) {

	v, err := protovalidate.New()
    if err != nil {
		return nil, models.ErrValidationFailed
    }

    if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}
	
    // TODO: получить id из jwt MW
	userId := "00000000-0000-0000-0000-0000000000000"

	form := &dto.UpdateUser{
		ID: dto.UserID(userId),
		Nickname: req.Nickname,
		Bio: req.Bio,
		Avatar: dto.Url(req.AvatarUrl),
	}

	res, err := s.UsersUsecase.UpdateUser(form)

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(res.ID),
		Nickname: res.Nickname,
		Bio: res.Bio,
		AvatarUrl: res.Avatar,
	}
	
	return &users.UpdateProfileResponse{UserProfile: userProfile}, nil
}

func (s *service) GetProfileByID(ctx context.Context, req *users.GetProfileByIDRequest) (*users.GetProfileByIDResponse, error) {

	res, err := s.UsersUsecase.GetProfileByID(dto.UserID(req.UserId))

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(res.ID),
		Nickname: res.Nickname,
		Bio: res.Bio,
		AvatarUrl: res.Avatar,
	}

	return &users.GetProfileByIDResponse{UserProfile: userProfile}, nil
}

func (s *service) GetProfileByNickname(ctx context.Context, req *users.GetProfileByNicknameRequest) (*users.GetProfileByNicknameResponse, error) {

	res, err := s.UsersUsecase.GetProfileByNickname(req.Nickname)

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(res.ID),
		Nickname: res.Nickname,
		Bio: res.Bio,
		AvatarUrl: res.Avatar,
	}

	return &users.GetProfileByNicknameResponse{UserProfile: userProfile}, nil
}

func (s *service) SearchByQuery(ctx context.Context, req *users.SearchByQueryRequest) (*users.SearchByQueryResponse, error) {

	return &users.SearchByQueryResponse{Profiles: []*users.UserProfile{&users.UserProfile{}}}, nil
}