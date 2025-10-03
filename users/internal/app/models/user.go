package models

type UserID uint64

type User struct {
	ID UserID
	Nickname string
	Bio string
	Avatar string
}