package usecases

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/auth/internal/app/utils"
)

func (ac *AuthUsecase) Login(ctx context.Context, a *dto.Login) (*models.Auth, error) {
	// Проверка наличия пользователя
	user, err := ac.RepoAuth.FetchByEmail(ctx, a.Email)

	if err != nil {
		return nil, ErrNotExistedUser
	}

	// Проверка правильности пароля
	hashedBytes := []byte(user.PasswordHash)
	passwordBytes := []byte(a.Password)

	if err := bcrypt.CompareHashAndPassword(hashedBytes, passwordBytes); err != nil {
		return nil, ErrBadCredentials
	}

	// Генерация jwt токенов
	authUser, err := utils.GenerateJWT(ctx, dto.UserID(user.ID))
	if err != nil {
		return nil, fmt.Errorf("Login: GenerateJWT: %w", err)
	}

	authRefresh := &dto.AuthRefresh{
		ID:           dto.UserID(authUser.ID),
		RefreshToken: authUser.Token.RefreshToken,
	}

	// Сохранение refresh токена в key-value базе
	if err := ac.RepoToken.Save(ctx, authRefresh); err != nil {
		return nil, fmt.Errorf("Login: RepoToken.Save: %w", err)
	}

	return authUser, nil
}
