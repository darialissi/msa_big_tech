package auth_grpc

import (
	"errors"
)

var (
	// Не прошла валидация запроса
	ErrValidationFailed = errors.New("Validation error")
)