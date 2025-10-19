package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/auth/internal/app/utils"
)

func (ac *AuthUsecase) Refresh(ctx context.Context, a *dto.AuthRefresh) (*models.Auth, error) {
	// Проверка переданного refresh токена в key-value базе
	if ref, err := ac.RepoToken.FetchById(ctx, models.UserID(a.ID)); err != nil {
		return nil, fmt.Errorf("Refresh: RepoToken.FetchById: %w", err)
	} else if ref != a.RefreshToken {
		return nil, InvalidRefreshToken
	}

	// Генерация jwt токенов
	authUser, err := utils.GenerateJWT(ctx, a.ID)
	if err != nil {
		return nil, fmt.Errorf("Refresh: GenerateJWT: %w", err)
	}

	authRefresh := &dto.AuthRefresh{
		ID:           a.ID,
		RefreshToken: authUser.Token.RefreshToken,
	}

	// Сохранение/перезапись refresh токена в key-value базе
	if err := ac.RepoToken.Save(ctx, authRefresh); err != nil {
		return nil, fmt.Errorf("Refresh: RepoToken.Save: %w", err)
	}

	return authUser, nil
}
