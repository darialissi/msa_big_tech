package friend

const friendsTable = "friends"

const (
	friendsTableColumnUserID    = "user_id"
	friendsTableColumnFriendID  = "friend_id"
	friendsTableColumnCreatedAt = "created_at"
)

var friendsTableColumns = []string{
	friendsTableColumnUserID,
	friendsTableColumnFriendID,
	friendsTableColumnCreatedAt,
}
