package users_grpc

import (
	"context"
	"fmt"

	"buf.build/go/protovalidate"

	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	users "github.com/darialissi/msa_big_tech/users/pkg"
)


func (s *server) CreateProfile(ctx context.Context, req *users.CreateProfileRequest) (*users.CreateProfileResponse, error) {

	v, err := protovalidate.New()
    if err != nil {
		return nil, models.ErrValidationFailed
    }

    if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

    userEmail, ok := ctx.Value("Email").(string) // TODO: прокинуть Email через middleware
    if !ok {
        return nil, fmt.Errorf("Email not found in context")
    }

	userInp := &dto.CreateUser{
		Email: userEmail,
		Nickname: req.Nickname,
		Bio: req.Bio,
		Avatar: dto.Url(req.AvatarUrl),
	}

	user, err := s.UsersUsecase.CreateUser(userInp)

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(user.ID),
		Email: user.Email,
		Nickname: user.Nickname,
		Bio: user.Bio,
		AvatarUrl: user.Avatar,
	}

	return &users.CreateProfileResponse{UserProfile: userProfile}, nil
}

func (s *server) UpdateProfile(ctx context.Context, req *users.UpdateProfileRequest) (*users.UpdateProfileResponse, error) {

	v, err := protovalidate.New()
    if err != nil {
		return nil, models.ErrValidationFailed
    }

    if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	userInp := &dto.UpdateUser{
		Nickname: req.Nickname,
		Bio: req.Bio,
		Avatar: dto.Url(req.AvatarUrl),
	}

	user, err := s.UsersUsecase.UpdateUser(userInp)

	if err != nil {
		return nil, err
	}

	userProfile := &users.UserProfile{
		Id: string(user.ID),
		Email: user.Email,
		Nickname: user.Nickname,
		Bio: user.Bio,
		AvatarUrl: user.Avatar,
	}
	
	return &users.UpdateProfileResponse{UserProfile: userProfile}, nil
}

func (s *server) GetProfileByID(ctx context.Context, req *users.GetProfileByIDRequest) (*users.GetProfileByIDResponse, error) {

	return &users.GetProfileByIDResponse{UserProfile: &users.UserProfile{}}, nil
}

func (s *server) GetProfileByNickname(ctx context.Context, req *users.GetProfileByNicknameRequest) (*users.GetProfileByNicknameResponse, error) {

	return &users.GetProfileByNicknameResponse{UserProfile: &users.UserProfile{Nickname: req.Nickname}}, nil
}

func (s *server) SearchByQuery(ctx context.Context, req *users.SearchByQueryRequest) (*users.SearchByQueryResponse, error) {

	return &users.SearchByQueryResponse{Profiles: []*users.UserProfile{&users.UserProfile{}}}, nil
}