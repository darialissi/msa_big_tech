package models

import (
	"time"
)

type ChatID string
type UserID string
type MessageID string

type DirectChat struct {
	ID        ChatID
	CreatorID UserID
	CreatedAt time.Time
}

type ChatMember struct {
	ChatID ChatID
	UserID UserID
}

type Message struct {
	ID        MessageID
	Text      string
	ChatID    ChatID
	SenderID  UserID
	CreatedAt time.Time
}

type Cursor struct {
	NextCursor string
	Limit      uint64
}
