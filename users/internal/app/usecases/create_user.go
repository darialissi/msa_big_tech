package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)


func (uc *UsersUsecase) CreateUser(u *dto.CreateUser) (*models.User, error) {
    // валидация, создание профиля пользователя

	if user, err := uc.repo.FetchByNickname(u.Nickname); err != nil {
		return nil, fmt.Errorf("CreateUser error: %w", err)
	} else if user != nil {
		return nil, ErrExistedNickname
	}

	userSave := &dto.SaveUser{
		Email: u.Email,
		Nickname: u.Nickname,
		Bio: u.Bio,
		Avatar: u.Avatar,
	}

	user, err := uc.repo.Save(userSave)
	if err != nil {
		return nil, fmt.Errorf("CreateUser error: %w", err)
	}

	return user, nil
}