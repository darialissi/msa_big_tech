package models

type UserID string

type User struct {
	ID UserID
	Nickname string
	Bio string
	Avatar string
}