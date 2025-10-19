package models

import (
	"time"
)

type UserID string

type User struct {
	ID        UserID
	Nickname  string
	Bio       string
	AvatarUrl string
	CreatedAt time.Time
}
