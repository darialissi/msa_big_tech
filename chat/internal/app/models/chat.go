package models

type ChatID string
type UserID string
type MessageID string


type DirectChat struct {
	ID ChatID
	CreatorID UserID
	ParticipantID UserID
	Messages []*Message
}

type ChatParticipant struct {
    ChatID   ChatID
    UserID   UserID
    Role     string // "member", "admin", "creator"
}

type Message struct {
	ID MessageID
	Text string
	ChatID ChatID
	FromUserID UserID
}