package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
)


func (uc *UsersUsecase) GetProfileByNickname(ctx context.Context, nickname string) (*models.User, error) {
    // получение профиля пользователя

	user, err := uc.repoUsers.FetchByNickname(ctx, nickname)

	if err != nil {
		return nil, fmt.Errorf("GetProfileByNickname: repoUsers.FetchByNickname: %w", err)
	}

	if user == nil {
		return nil, ErrNoProfileFound
	}

	return user, nil
}