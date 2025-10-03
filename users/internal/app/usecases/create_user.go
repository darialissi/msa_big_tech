package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)


func (uc *UsersUsecase) CreateUser(u *dto.CreateUser) (*models.User, error) {
    // валидация, создание профиля пользователя

	if user, err := uc.repo.FetchByNickname(u.Nickname); err != nil {
		return nil, fmt.Errorf("CreateUser: FetchByNickname: %w", err)
	} else if user != nil {
		return nil, ErrExistedNickname
	}

	user, err := uc.repo.Save(u)
	if err != nil {
		return nil, fmt.Errorf("CreateUser: Save: %w", err)
	}

	return user, nil
}