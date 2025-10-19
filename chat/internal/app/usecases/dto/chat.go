package dto

type ChatID string
type UserID string
type MessageID string

type CreateDirectChat struct {
	CreatorID UserID
	MemberID  UserID
}

type SendMessage struct {
	Text     string
	SenderID UserID
	ChatID   ChatID
}

type ListMessages struct {
	ChatID ChatID
	Limit  uint64
	Cursor MessageID
}

type StreamMessages struct {
	ChatID    ChatID
	SinceUnix int64
}
