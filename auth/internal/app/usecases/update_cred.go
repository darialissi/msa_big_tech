package usecases

import (
	"fmt"

  	"golang.org/x/crypto/bcrypt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)


func (ac *AuthUsecase) UpdateCred(a *dto.UpdateCred) (*models.User, error) {
    // валидация, проверка уникальности email, обновление кредов в бд
	if !a.Password.IsValid() {
		return nil, ErrWeakPassword
	}

	user, err := ac.RepoAuth.FetchByEmail(a.Email)

	if err != nil  {
		return nil, fmt.Errorf("UpdateCred error: %w", err)
	} else if user != nil && dto.UserID(user.ID) != a.ID {
		return nil, ErrRegisteredEmail
	}

	passwordBytes := []byte(a.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return nil, fmt.Errorf("UpdateCred error: %w", err)
	}

	userInp := &dto.UpdateCredSave{
		ID: dto.UserID(user.ID),
		Email: a.Email,
		PasswordHash: string(hashedPasswordBytes),
	}

	if _, err := ac.RepoAuth.Update(userInp); err != nil {
		return nil, fmt.Errorf("UpdateCred error: %w", err)
	}

	return user, nil
}