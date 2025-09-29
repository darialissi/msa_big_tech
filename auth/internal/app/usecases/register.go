package usecases

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)


func (ac *AuthUsecase) Register(a *dto.Register) (*models.User, error) {
    // валидация, проверка уникальности данных, хеширование пароля, запись в бд
	if !a.Password.IsValid() {
		return nil, ErrWeakPassword
	}

	if existed, err := ac.repo.FetchByEmail(a.Email); err != nil  {
		return nil, fmt.Errorf("Register error: %w", err)
	} else if existed != nil {
		return nil, ErrRegisteredEmail
	}

	passwordBytes := []byte(a.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return nil, fmt.Errorf("Register error: %w", err)
	}

	credInp := &dto.SaveCred{
		Email: a.Email,
		PasswordHash: string(hashedPasswordBytes),
	}

	user, err := ac.repo.Save(credInp)

	if err != nil {
		return nil, fmt.Errorf("Register error: %w", err)
	}

	return user, nil
}