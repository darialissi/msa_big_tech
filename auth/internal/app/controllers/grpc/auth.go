package auth_grpc

import (
	"context"

	"buf.build/go/protovalidate"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
	auth "github.com/darialissi/msa_big_tech/auth/pkg"
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
		Email:    req.Email,
		Password: dto.Password(req.Password),
	}

	userResponse, err := s.AuthUsecase.Register(ctx, form)

	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{UserId: string(userResponse.ID)}, nil
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
		Email:    req.Email,
		Password: dto.Password(req.Password),
	}

	authUser, err := s.AuthUsecase.Login(ctx, form)

	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		UserId:       string(authUser.ID),
		AccessToken:  authUser.Token.AccessToken,
		RefreshToken: authUser.Token.RefreshToken,
	}, nil
}

func (s *service) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {

	v, err := protovalidate.New()
	if err != nil {
		return nil, models.ErrValidationFailed
	}

	if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	// TODO: получить id из jwt MW
	userId := uuid.New().String()

	form := &dto.AuthRefresh{
		ID:           dto.UserID(userId),
		RefreshToken: req.RefreshToken,
	}

	authUser, err := s.AuthUsecase.Refresh(ctx, form)

	if err != nil {
		return nil, err
	}

	return &auth.RefreshResponse{
		UserId:       string(authUser.ID),
		AccessToken:  authUser.Token.AccessToken,
		RefreshToken: authUser.Token.RefreshToken,
	}, nil
}
