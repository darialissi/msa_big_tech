package models

type ChatID uint64
type UserID uint64
type MessageID uint64


type Chat struct {
	ID ChatID
	Name string
	CreatorID UserID
	ParticipantIDs []UserID
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