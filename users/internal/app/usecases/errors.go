package usecases

import (
	"errors"
)

// Ошибки уровня бизнес логики
var (
	// Nickname уже существует
	ErrExistedNickname = errors.New("Nickname is already exists")
	// Пользователь не существует
	ErrNotExisted = errors.New("User does not exist")
)