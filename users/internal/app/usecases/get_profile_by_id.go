package usecases

import (
	"fmt"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)


func (uc *UsersUsecase) GetProfileByID(id dto.UserID) (*models.User, error) {
    // получение профиля пользователя

	user, err := uc.repo.FetchById(id)

	if err != nil {
		return nil, fmt.Errorf("GetProfileByID: FetchById: %w", err)
	}

	if user == nil {
		return nil, ErrProfileEmpty
	}

	return user, nil
}