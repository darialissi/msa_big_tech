package models

import (
	"strings"
    "time"
)

type UserID string
type FriendRequestID string
type FriendRequestStatus int64


type FriendRequest struct {
	ID FriendRequestID
	Status FriendRequestStatus
	FromUserID UserID
	ToUserID UserID
    CreatedAt time.Time
}

type UserFriend struct {
	UserID UserID
	FriendID UserID
    CreatedAt time.Time
}

type Cursor struct {
	NextCursor string
	Limit uint64
}

const (
	// FriendRequestStatusPending заявка отправлена, ожидает ответа
	FriendRequestStatusPending FriendRequestStatus = iota + 1
	// FriendRequestStatusAccepted заявка принята
	FriendRequestStatusAccepted
	// FriendRequestStatusDeclined заявка отклонена
	FriendRequestStatusDeclined
)

// String возвращает строковое представление статуса
func (s FriendRequestStatus) String() string {
	switch s {
	case FriendRequestStatusPending:
		return "PENDING"
	case FriendRequestStatusAccepted:
		return "ACCEPTED"
	case FriendRequestStatusDeclined:
		return "DECLINED"
	default:
		return "UNKNOWN"
	}
}

// ParseFriendRequestStatus преобразует строку в FriendRequestStatus
func ParseFriendRequestStatus(status string) (FriendRequestStatus, error) {
	switch strings.ToUpper(status) {
	case "PENDING":
		return FriendRequestStatusPending, nil
	case "ACCEPTED":
		return FriendRequestStatusAccepted, nil
	case "DECLINED":
		return FriendRequestStatusDeclined, nil
	default:
		return 0, ErrInvalidFriendRequestStatus
	}
}