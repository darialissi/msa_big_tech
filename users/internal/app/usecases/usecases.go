package usecases

import (
	"context"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)

type UsersUsecases interface {
	// Создание профиля пользователя
	CreateUser(ctx context.Context, u *dto.CreateUser) (*models.User, error)
	// Обновление профиля пользователя
	UpdateUser(ctx context.Context, u *dto.UpdateUser) (*models.User, error)
	// Получение профиля пользователя по id
	GetProfileByID(ctx context.Context, id dto.UserID) (*models.User, error)
	// Получение профиля пользователя по nickname
	GetProfileByNickname(ctx context.Context, nickname string) (*models.User, error)
}

type UsersRepository interface {
	Save(ctx context.Context, in *models.User) (*models.User, error)
	Update(ctx context.Context, in *models.User) (*models.User, error)
	FetchById(ctx context.Context, id models.UserID) (*models.User, error)
	FetchByNickname(ctx context.Context, nickname string) (*models.User, error)
}

// Проверка реализации всех методов интерфейса при компиляции
var _ UsersUsecases = (*UsersUsecase)(nil)

type TxManager interface {
	RunReadCommitted(ctx context.Context, f func(txCtx context.Context) error) error
}

type UsersUsecase struct {
	repoUsers UsersRepository
	txMan     TxManager
}

func NewUsersUsecase(repo UsersRepository, txMan TxManager) *UsersUsecase {
	return &UsersUsecase{
		repoUsers: repo,
		txMan:     txMan,
	}
}
