package dto

type ChatID uint64
type UserID uint64
type MessageID uint64


type CreateChat struct {
	Name string
	CreatorID UserID
	ParticipantIDs []UserID
}

type SendMessage struct {
	Text string
	FromUserID UserID
	ToChatID ChatID
}