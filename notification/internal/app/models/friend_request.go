package models

import (
	"time"
)

// FriendRequest - запрос в друзья
type FriendRequest struct {
	ID        string    `json:"id"`
	FromUser  string    `json:"from_user"`
	ToUser    string    `json:"to_user"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
