package usecases

import (
	"errors"
)

// Ошибки уровня бизнес логики
var (
	// Nickname уже существует
	ErrExistedNickname = errors.New("Nickname is already exists")
	// Профиль пользователя не создан
	ErrProfileEmpty = errors.New("User Profile is empty")
)