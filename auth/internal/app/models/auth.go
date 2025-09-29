package models

type UserID uint64

type User struct {
	ID UserID
	Email string
    PasswordHash string
}

type Token struct {
	AccessToken string
	RefreshToken string
}

type Auth struct {
	ID UserID
	Token Token
}