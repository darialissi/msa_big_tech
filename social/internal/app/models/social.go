package models

import (
	"strings"
)

type UserID uint64
type FriendID uint64
type FriendRequestID uint64
type FriendRequestStatus uint64


type FriendRequest struct {
	ID FriendRequestID
	Status FriendRequestStatus
	FromUserID UserID
	ToUserID UserID
}

type UserFriend struct {
	ID uint64
	UserId UserID
	FriendId UserID
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