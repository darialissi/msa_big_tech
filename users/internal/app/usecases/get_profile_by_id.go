package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)

func (uc *UsersUsecase) GetProfileByID(ctx context.Context, id dto.UserID) (*models.User, error) {
	// получение профиля пользователя

	user, err := uc.repoUsers.FetchById(ctx, models.UserID(id))

	if err != nil {
		return nil, fmt.Errorf("GetProfileByID: repoUsers.FetchById: %w", err)
	}

	if user == nil {
		return nil, ErrNoProfileFound
	}

	return user, nil
}
