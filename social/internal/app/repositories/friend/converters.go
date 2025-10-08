package friend

import (
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)

// FriendRequestRow — «плоская» проекция строки таблицы friend_requests
// Nullable поля представлены sql.Null*.
type FriendRow struct {
	UserID  	uuid.UUID		`db:"user_id"`
	FriendID    uuid.UUID		`db:"friend_id"`
	CreatedAt 	time.Time       `db:"created_at"`
}

func (row *FriendRow) Values() []any {
	return []any{
		row.UserID, row.FriendID, row.CreatedAt,
	}
}

// ToModel конвертирует FriendRow (sql.Null*) в доменную модель models.UserFriend (*string/*time.Time).
func ToModel(r *FriendRow) *models.UserFriend {
	if r == nil {
		return nil
	}
	return &models.UserFriend{
		UserID: models.UserID(r.UserID.String()),
		FriendID: models.UserID(r.FriendID.String()),
		CreatedAt: r.CreatedAt,

	}
}

// FromSaveDTO конвертирует dto в FriendRow для INSERT
func FromSaveDTO(d *dto.SaveFriend) FriendRow {
	if d == nil {
		return FriendRow{}
	}

	parsedUserID, err := uuid.Parse(string(d.UserID))
	if err != nil {
		return FriendRow{}
	}

	parsedFriendID, err := uuid.Parse(string(d.FriendID))
	if err != nil {
		return FriendRow{}
	}

	return FriendRow{
		UserID: parsedUserID,
		FriendID: parsedFriendID,
	}
}

// FromDeleteDTO конвертирует доменную модель в FriendRow для DELETE
func FromDeleteDTO(d *dto.RemoveFriend) FriendRow {
	if d == nil {
		return FriendRow{}
	}

	parsedUserID, err := uuid.Parse(string(d.UserID))
	if err != nil {
		return FriendRow{}
	}

	parsedFriendID, err := uuid.Parse(string(d.FriendID))
	if err != nil {
		return FriendRow{}
	}

	return FriendRow{
		UserID: parsedUserID,
		FriendID: parsedFriendID,
	}
}