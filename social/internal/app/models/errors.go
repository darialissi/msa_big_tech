package models

import (
	"errors"
)

var (
	// Не прошла валидация запроса
	ErrValidationFailed           = errors.New("Validation error")
	ErrInvalidFriendRequestStatus = errors.New("Invalid Friend Request status")
)
