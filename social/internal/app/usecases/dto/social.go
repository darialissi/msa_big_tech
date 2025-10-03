package dto

type UserID uint64
type FriendID uint64
type FriendRequestID uint64
type FriendRequestStatus uint64


type SendFriendRequest struct {
	FromUserID UserID
	ToUserID UserID
}

type SaveFriendRequest struct {
	StatusID FriendRequestStatus
	FromUserID UserID
	ToUserID UserID
}

type ChangeStatus struct {
	ReqID FriendRequestID
	StatusID FriendRequestStatus
}

type SaveFriend struct {
	UserID UserID
	FriendID FriendID
}

type RemoveFriend struct {
	UserID UserID
	FriendID FriendID
}