package usecases

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/auth/internal/app/utils"
)


func (ac *AuthUsecase) Login(a *dto.Login) (*models.Auth, error) {
	// проверка наличия пользователя, проверка пароля, генерация токенов и запись refresh
	user, err := ac.RepoAuth.FetchByEmail(a.Email)

	if err != nil {
		return nil, ErrNotExistedUser
	}

	hashedBytes := []byte(user.PasswordHash)
	passwordBytes := []byte(a.Password)
	
	if err := bcrypt.CompareHashAndPassword(hashedBytes, passwordBytes); err != nil {
		return nil, ErrBadCredentials
	}

	authUser, err := utils.GenerateJWT(dto.UserID(user.ID))
	if err != nil {
		return nil, fmt.Errorf("Login error: %w", err)
	}

	authRefresh := &dto.AuthRefresh{
		ID: dto.UserID(authUser.ID),
		RefreshToken: authUser.Token.RefreshToken,
	}

	if err := ac.RepoToken.Save(authRefresh); err != nil {
		return nil, fmt.Errorf("Login error: %w", err)
	}

	return authUser, nil
}