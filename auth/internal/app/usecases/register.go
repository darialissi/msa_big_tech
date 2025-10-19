package usecases

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)

func (ac *AuthUsecase) Register(ctx context.Context, a *dto.Register) (*models.User, error) {
	// Валидация пароля
	if !a.Password.IsValid() {
		return nil, ErrWeakPassword
	}

	// Проверка уникальности email
	if existed, err := ac.RepoAuth.FetchByEmail(ctx, a.Email); err != nil {
		return nil, fmt.Errorf("Register: RepoAuth.FetchByEmail: %w", err)
	} else if existed != nil {
		return nil, ErrRegisteredEmail
	}

	// Хэширование пароля
	passwordBytes := []byte(a.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return nil, fmt.Errorf("Register: GenerateFromPassword: %w", err)
	}

	model := &models.User{
		Email:        a.Email,
		PasswordHash: string(hashedPasswordBytes),
	}

	// Сохранение данных пользователя
	user, err := ac.RepoAuth.Save(ctx, model)

	if err != nil {
		return nil, fmt.Errorf("Register: RepoAuth.Save: %w", err)
	}

	return user, nil
}
