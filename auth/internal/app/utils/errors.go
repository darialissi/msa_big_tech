package utils

import (
	"errors"
)

var (
	// Не найден JWT_SECRET в .env
	ErrSecretNotFound = errors.New("JWT_SECRET does not exist!")
	// Генерация Access токена
	ErrAccessToken = errors.New("Access Token generate error")
	// Генерация Refresh токена
	ErrRefreshToken = errors.New("Refresh Token generate error")
)
