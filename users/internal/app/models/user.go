package models

type UserID uint64

type User struct {
	ID UserID
	Email string
	Nickname string
	Bio string
	Avatar string
}