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

	if existed, err := ac.RepoAuth.FetchByEmail(a.Email); err != nil  {
		return nil, fmt.Errorf("Register: FetchByEmail: %w", err)
	} else if existed != nil {
		return nil, ErrRegisteredEmail
	}

	passwordBytes := []byte(a.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return nil, fmt.Errorf("Register: GenerateFromPassword: %w", err)
	}

	credInp := &dto.SaveCred{
		Email: a.Email,
		PasswordHash: string(hashedPasswordBytes),
	}

	user, err := ac.RepoAuth.Save(credInp)

	if err != nil {
		return nil, fmt.Errorf("Register: Save: %w", err)
	}

	return user, nil
}