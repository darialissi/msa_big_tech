package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
)


func (uc *UsersUsecase) GetProfileByNickname(nickname string) (*models.User, error) {
    // получение профиля пользователя

	user, err := uc.repo.FetchByNickname(nickname)

	if err != nil {
		return nil, fmt.Errorf("GetProfileByNickname: FetchByNickname: %w", err)
	}

	if user == nil {
		return nil, ErrProfileEmpty
	}

	return user, nil
}