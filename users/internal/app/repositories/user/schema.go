package user


const usersTable = "users"

const (
	usersTableColumnID     		= "id"
	usersTableColumnNickname    = "nickname"
	usersTableColumnAvatarUrl   = "avatar_url"
	usersTableColumnBio     	= "bio"
	usersTableColumnCreatedAt 	= "created_at"
)

var usersTableColumns = []string{
	usersTableColumnID,
	usersTableColumnNickname,
	usersTableColumnAvatarUrl,
	usersTableColumnBio,
	usersTableColumnCreatedAt,
}