package usecases

import (
	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)


type AuthUsecases interface {
	// Регистрация
    Register(a *dto.Register) (*models.User, error)
	// Логин
    Login(a *dto.Login) (*models.Auth, error)
	// Обновление токенов
    Refresh(a *dto.AuthRefresh) (*models.Auth, error)
	// Обновление email/password
    UpdateCred(a *dto.UpdateCred) (*models.User, error)
}

type AuthRepository interface {
    Save(user *dto.SaveCred) (*models.User, error)
    Update(user *dto.UpdateCredSave) (*models.User, error)
    FetchById(userId dto.UserID) (*models.User, error)
    FetchByEmail(email string) (*models.User, error)
}

type TokenRepository interface {
	Save(auth *dto.AuthRefresh) (error)
    Fetch(userId dto.UserID) (string, error)
}

// Проверка реализации всех методов интерфейса при компиляции
var _ AuthUsecases = (*AuthUsecase)(nil)

type AuthUsecase struct {
    repo AuthRepository
	repoToken TokenRepository
}

func NewAuthUsecase(repo AuthRepository, repoToken TokenRepository) *AuthUsecase {
    return &AuthUsecase{
        repo:      repo,
        repoToken: repoToken,
    }
}