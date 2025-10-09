package dto


type UserID string
type FriendRequestID string
type FriendRequestStatus int64


type SendFriendRequest struct {
	FromUserID UserID
	ToUserID UserID
}

type RemoveFriend struct {
	UserID UserID
	FriendID UserID
}