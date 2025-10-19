package chat

const chatsTable = "chats"

const (
	chatsTableColumnID        = "id"
	chatsTableColumnCreatorID = "creator_id"
	chatsTableColumnCreatedAt = "created_at"
)

var chatsTableColumns = []string{
	chatsTableColumnID,
	chatsTableColumnCreatorID,
	chatsTableColumnCreatedAt,
}
