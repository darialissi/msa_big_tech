package dto


type UserID string
type FriendRequestID string
type FriendRequestStatus int64


type SendFriendRequest struct {
	FromUserID UserID
	ToUserID UserID
}

type SaveFriendRequest struct {
	Status FriendRequestStatus
	FromUserID UserID
	ToUserID UserID
}

type UpdateFriendRequest struct {
	ReqID FriendRequestID
	Status FriendRequestStatus
}

type SaveFriend struct {
	UserID UserID
	FriendID UserID
}

type RemoveFriend struct {
	UserID UserID
	FriendID UserID
}