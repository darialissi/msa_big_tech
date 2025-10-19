package usecases

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)

func (uc *UsersUsecase) UpdateUser(ctx context.Context, u *dto.UpdateUser) (*models.User, error) {
	// обновление профиля пользователя

	userProfile, err := uc.repoUsers.FetchById(ctx, models.UserID(u.ID))

	if err != nil {
		return nil, fmt.Errorf("UpdateUser: repoUsers.FetchById: %w", err)
	}

	if userProfile.Nickname != u.Nickname {
		if existed, err := uc.repoUsers.FetchByNickname(ctx, u.Nickname); err != nil {
			return nil, fmt.Errorf("UpdateUser: repoUsers.FetchByNickname: %w", err)
		} else if existed != nil {
			return nil, ErrExistedNickname
		}
	}

	model := &models.User{
		ID:        models.UserID(u.ID),
		Nickname:  u.Nickname,
		AvatarUrl: string(u.AvatarUrl),
		Bio:       u.Bio,
	}

	user, err := uc.repoUsers.Update(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("UpdateUser: repoUsers.Update: %w", err)
	}

	return user, nil
}
