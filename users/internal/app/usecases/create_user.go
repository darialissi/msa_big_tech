package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)


func (uc *UsersUsecase) CreateUser(ctx context.Context, u *dto.CreateUser) (*models.User, error) {
    // валидация, создание профиля пользователя

	if user, err := uc.repoUsers.FetchByNickname(ctx, u.Nickname); err != nil {
		return nil, fmt.Errorf("CreateUser: repoUsers.FetchByNickname: %w", err)
	} else if user != nil {
		return nil, ErrExistedNickname
	}

	model := &models.User{
		ID: models.UserID(u.ID),
		Nickname: u.Nickname,
		AvatarUrl: string(u.AvatarUrl),
		Bio: u.Bio,
	}

	user, err := uc.repoUsers.Save(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("CreateUser: repoUsers.Save: %w", err)
	}

	return user, nil
}