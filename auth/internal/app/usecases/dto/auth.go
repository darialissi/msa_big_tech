package dto

import (
	"strconv"
	"unicode"
)


type Password string
type UserID uint64


type Register struct {
	Email string
	Password Password

}

type Login struct {
	Email string
	Password Password
}

type SaveCred struct {
	Email string
    PasswordHash string
}

type UpdateCred struct {
	ID UserID
	Email string
    Password Password
}

type UpdateCredSave struct {
	ID UserID
	Email string
    PasswordHash string
}

type AuthRefresh struct {
	ID UserID
	RefreshToken string
}

func (id *UserID) String() string {
	str := strconv.Itoa(int(*id))

	return str
}

func (password Password) IsValid() bool {
	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, r := range password {
		switch {
			case unicode.IsUpper(r): hasUpper = true
			case unicode.IsLower(r): hasLower = true
			case unicode.IsDigit(r): hasUpper = true
		}
	}

	return hasUpper && hasLower && hasDigit
}