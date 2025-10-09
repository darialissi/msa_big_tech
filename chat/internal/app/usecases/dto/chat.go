package dto

type ChatID string
type UserID string
type MessageID string


type CreateDirectChat struct {
	CreatorID UserID
	MemberID UserID
}

type SendMessage struct {
	Text string
	SenderID UserID
	ChatID ChatID
}