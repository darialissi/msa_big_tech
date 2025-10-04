package dto

type ChatID string
type UserID string
type MessageID string


type CreateDirectChat struct {
	CreatorID UserID
	ParticipantID UserID
}

type SendMessage struct {
	Text string
	FromUserID UserID
	ToChatID ChatID
}