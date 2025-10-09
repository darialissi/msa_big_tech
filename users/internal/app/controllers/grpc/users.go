package users_grpc

import (
	"context"

	"buf.build/go/protovalidate"

	"github.com/google/uuid"

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
	userId := uuid.New().String()

	form := &dto.CreateUser{
		ID: dto.UserID(userId),
		Nickname: req.Nickname,
		Bio: req.Bio,
		AvatarUrl: dto.Url(req.AvatarUrl),
	}

	res, err := s.UsersUsecase.CreateUser(ctx, form)

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(res.ID),
		Nickname: res.Nickname,
		Bio: res.Bio,
		AvatarUrl: res.AvatarUrl,
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
	userId := uuid.New().String()

	form := &dto.UpdateUser{
		ID: dto.UserID(userId),
		Nickname: req.Nickname,
		Bio: req.Bio,
		AvatarUrl: dto.Url(req.AvatarUrl),
	}

	res, err := s.UsersUsecase.UpdateUser(ctx, form)

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(res.ID),
		Nickname: res.Nickname,
		Bio: res.Bio,
		AvatarUrl: res.AvatarUrl,
	}
	
	return &users.UpdateProfileResponse{UserProfile: userProfile}, nil
}

func (s *service) GetProfileByID(ctx context.Context, req *users.GetProfileByIDRequest) (*users.GetProfileByIDResponse, error) {

	res, err := s.UsersUsecase.GetProfileByID(ctx, dto.UserID(req.UserId))

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(res.ID),
		Nickname: res.Nickname,
		Bio: res.Bio,
		AvatarUrl: res.AvatarUrl,
	}

	return &users.GetProfileByIDResponse{UserProfile: userProfile}, nil
}

func (s *service) GetProfileByNickname(ctx context.Context, req *users.GetProfileByNicknameRequest) (*users.GetProfileByNicknameResponse, error) {

	res, err := s.UsersUsecase.GetProfileByNickname(ctx, req.Nickname)

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(res.ID),
		Nickname: res.Nickname,
		Bio: res.Bio,
		AvatarUrl: res.AvatarUrl,
	}

	return &users.GetProfileByNicknameResponse{UserProfile: userProfile}, nil
}

func (s *service) SearchByQuery(ctx context.Context, req *users.SearchByQueryRequest) (*users.SearchByQueryResponse, error) {

	return &users.SearchByQueryResponse{Profiles: []*users.UserProfile{&users.UserProfile{}}}, nil
}