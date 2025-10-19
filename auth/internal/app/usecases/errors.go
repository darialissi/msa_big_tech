package usecases

import (
	"errors"
)

// Ошибки уровня бизнес логики
var (
	// Невалидный Password
	ErrWeakPassword = errors.New("Password must contain at least 5 symbols with lower|upper|digit")
	// Неверный Email/Password
	ErrBadCredentials = errors.New("Email or Password is incorrect")
	// Email уже зарегистрирован
	ErrRegisteredEmail = errors.New("User with this Email is already registered")
	// Пользователь не существует
	ErrNotExistedUser = errors.New("User does not exist")
	// Невалидный refresh токен
	InvalidRefreshToken = errors.New("Refresh token is invalid")
)
