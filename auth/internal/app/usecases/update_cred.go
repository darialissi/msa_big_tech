package usecases

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)

func (ac *AuthUsecase) UpdateCred(ctx context.Context, a *dto.UpdateCred) (*models.User, error) {
	// Валидация пароля
	if !a.Password.IsValid() {
		return nil, ErrWeakPassword
	}

	// Проверка наличия пользователя
	user, err := ac.RepoAuth.FetchById(ctx, models.UserID(a.ID))

	if err != nil {
		return nil, fmt.Errorf("UpdateCred: RepoAuth.FetchById %w", err)
	}

	// Проверка уникальности email (если были изменения)
	if user.Email != a.Email {
		if existed, err := ac.RepoAuth.FetchByEmail(ctx, a.Email); err != nil {
			return nil, fmt.Errorf("UpdateUser: RepoAuth.FetchByEmail: %w", err)
		} else if existed != nil {
			return nil, ErrRegisteredEmail
		}
	}

	// Хэширование пароля
	passwordBytes := []byte(a.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return nil, fmt.Errorf("UpdateCred: GenerateFromPassword: %w", err)
	}

	model := &models.User{
		ID:           user.ID,
		Email:        a.Email,
		PasswordHash: string(hashedPasswordBytes),
	}

	// Обновление данных пользователя
	if _, err := ac.RepoAuth.Update(ctx, model); err != nil {
		return nil, fmt.Errorf("UpdateCred: Update: %w", err)
	}

	return user, nil
}
