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

type ListFriends struct {
	UserID UserID
	Limit uint64
	Cursor UserID
}