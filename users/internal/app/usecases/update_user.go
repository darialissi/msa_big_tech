package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)


func (uc *UsersUsecase) UpdateUser(u *dto.UpdateUser) (*models.User, error) {
    // обновление профиля пользователя

	if user, err := uc.repo.FetchByNickname(u.Nickname); err != nil {
		return nil, fmt.Errorf("UpdateUser error: %w", err)
	} else if user != nil && dto.UserID(user.ID) != u.ID {
		return nil, ErrExistedNickname
	}

	user, err := uc.repo.Update(u)
	if err != nil {
		return nil, fmt.Errorf("UpdateUser error: %w", err)
	}

	return user, nil
}