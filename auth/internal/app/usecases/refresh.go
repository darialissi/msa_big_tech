package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/auth/internal/app/utils"
)


func (ac *AuthUsecase) Refresh(a *dto.AuthRefresh) (*models.Auth, error) {
	// обновление токенов, запись в хранилище

	authUser, err := utils.GenerateJWT(a.ID)
	if err != nil {
		return nil, fmt.Errorf("Refresh: GenerateJWT: %w", err)
	}

	authRefresh := &dto.AuthRefresh{
		ID: a.ID,
		RefreshToken: authUser.Token.RefreshToken,
	}

	if err := ac.RepoToken.Save(authRefresh); err != nil {
		return nil, fmt.Errorf("Refresh: Save: %w", err)
	}

	return authUser, nil
}