package dto

import (
	"unicode"
)

type Password string
type UserID string

type Register struct {
	Email    string
	Password Password
}

type Login struct {
	Email    string
	Password Password
}

type UpdateCred struct {
	ID       UserID
	Email    string
	Password Password
}

type AuthRefresh struct {
	ID           UserID
	RefreshToken string
}

func (password Password) IsValid() bool {
	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}
