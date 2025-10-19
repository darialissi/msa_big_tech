package usecases

import (
	"context"
	"errors"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)

type AuthUsecases interface {
	// Регистрация
	Register(ctx context.Context, a *dto.Register) (*models.User, error)
	// Логин
	Login(ctx context.Context, a *dto.Login) (*models.Auth, error)
	// Обновление токенов
	Refresh(ctx context.Context, a *dto.AuthRefresh) (*models.Auth, error)
	// Обновление email/password
	UpdateCred(ctx context.Context, a *dto.UpdateCred) (*models.User, error)
}

type AuthRepository interface {
	Save(ctx context.Context, in *models.User) (*models.User, error)
	Update(ctx context.Context, in *models.User) (*models.User, error)
	FetchById(ctx context.Context, userId models.UserID) (*models.User, error)
	FetchByEmail(ctx context.Context, email string) (*models.User, error)
}

type TokenRepository interface {
	Save(ctx context.Context, auth *dto.AuthRefresh) error
	FetchById(ctx context.Context, userId models.UserID) (string, error)
}

// Проверка реализации всех методов интерфейса при компиляции
var _ AuthUsecases = (*AuthUsecase)(nil)

type TxManager interface {
	RunReadCommitted(ctx context.Context, f func(txCtx context.Context) error) error
}

type Deps struct {
	RepoAuth  AuthRepository
	RepoToken TokenRepository
	TxMan     TxManager
}

type AuthUsecase struct {
	Deps // встраивание
}

func NewAuthUsecase(deps Deps) (*AuthUsecase, error) {
	if deps.RepoAuth == nil {
		return nil, errors.New("AuthRepository is required")
	}
	if deps.RepoToken == nil {
		return nil, errors.New("TokenRepository is required")
	}
	if deps.TxMan == nil {
		return nil, errors.New("TxMan is required")
	}

	return &AuthUsecase{deps}, nil
}
