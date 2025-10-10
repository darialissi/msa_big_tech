package auth


const authUsersTable = "auth_users"

const (
	authUsersTableColumnID     			= "id"
	authUsersTableColumnEmail     		= "email"
	authUsersTableColumnPasswordHash    = "password_hash"
	authUsersTableColumnCreatedAt 		= "created_at"
)

var authUsersTableColumns = []string{
	authUsersTableColumnID,
	authUsersTableColumnEmail,
	authUsersTableColumnPasswordHash,
	authUsersTableColumnCreatedAt,
}