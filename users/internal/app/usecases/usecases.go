package usecases

import (
	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)

type UsersUsecases interface {
	// Создание профиля пользователя
	CreateUser(u *dto.CreateUser) (*models.User, error)
	// Обновление профиля пользователя
	UpdateUser(u *dto.UpdateUser) (*models.User, error)
	// Получение профиля пользователя по id
	GetProfileByID(id dto.UserID) (*models.User, error)
	// Получение профиля пользователя по nickname
	GetProfileByNickname(nickname string) (*models.User, error)
}

type UsersRepository interface {
    Save(user *dto.CreateUser) (*models.User, error)
    Update(user *dto.UpdateUser) (*models.User, error)
    FetchById(id dto.UserID) (*models.User, error)
    FetchByNickname(nickname string) (*models.User, error)
}

// Проверка реализации всех методов интерфейса при компиляции
var _ UsersUsecases = (*UsersUsecase)(nil)

type UsersUsecase struct {
    repo UsersRepository
}

func NewUsersUsecase(repo UsersRepository) *UsersUsecase {
    return &UsersUsecase{
        repo: repo,
    }
}