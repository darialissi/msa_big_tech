package models

import (
	"time"
)

type UserID string

type User struct {
	ID UserID
	Email string
    PasswordHash string
	CreatedAt time.Time
}

type Token struct {
	AccessToken string
	RefreshToken string
}

type Auth struct {
	ID UserID
	Token Token
}