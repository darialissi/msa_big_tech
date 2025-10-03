package auth_grpc

import (
	"context"

	"buf.build/go/protovalidate"

	auth "github.com/darialissi/msa_big_tech/auth/pkg"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
)


func (s *service) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {

	v, err := protovalidate.New()
    if err != nil {
		return nil, models.ErrValidationFailed
    }

    if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	form := &dto.Register{
		Email: req.Email,
		Password: dto.Password(req.Password),
	}

	userResponse, err := s.AuthUsecase.Register(form)

	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{UserId: uint64(userResponse.ID)}, nil
}

func (s *service) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {

	v, err := protovalidate.New()
    if err != nil {
		return nil, models.ErrValidationFailed
    }

    if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	form := &dto.Login{
		Email: req.Email,
		Password: dto.Password(req.Password),
	}

	authUser, err := s.AuthUsecase.Login(form)

	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		UserId: uint64(authUser.ID), 
		AccessToken: authUser.Token.AccessToken, 
		RefreshToken: authUser.Token.RefreshToken,
		}, nil
}

func (s *service) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	form := &dto.AuthRefresh{
		ID: dto.UserID(req.UserId),
		RefreshToken: req.RefreshToken,
	}

	authUser, err := s.AuthUsecase.Refresh(form)

	if err != nil {
		return nil, err
	}

	return &auth.RefreshResponse{
		UserId: uint64(authUser.ID), 
		AccessToken: authUser.Token.AccessToken, 
		RefreshToken: authUser.Token.RefreshToken,
		}, nil
}
