package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
)

func (uc *UsersUsecase) GetProfileByNickname(ctx context.Context, nickname string) (*models.User, error) {

	// Получение профиля пользователя по Nickname
	user, err := uc.repoUsers.FetchByNickname(ctx, nickname)

	if err != nil {
		return nil, fmt.Errorf("GetProfileByNickname: repoUsers.FetchByNickname: %w", err)
	}

	if user == nil {
		return nil, ErrNoProfileFound
	}

	return user, nil
}
