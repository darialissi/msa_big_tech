package friend_request

import (
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)

// FriendRequestRow — «плоская» проекция строки таблицы friend_requests
// Nullable поля представлены sql.Null*.
type FriendRequestRow struct {
	ID        	uuid.UUID		`db:"id"`
	FromUserID  uuid.UUID		`db:"from_user_id"`
	ToUserID    uuid.UUID		`db:"to_user_id"`
	Status      uint64          `db:"status"`
	CreatedAt 	time.Time       `db:"created_at"`
}

func (row *FriendRequestRow) Values() []any {
	return []any{
		row.ID, row.FromUserID, row.ToUserID, row.Status, row.CreatedAt,
	}
}

// ToModel конвертирует FriendRequestRow (sql.Null*) в доменную модель models.FriendRequest (*string/*time.Time).
func ToModel(r *FriendRequestRow) *models.FriendRequest {
	if r == nil {
		return nil
	}
	return &models.FriendRequest{
		ID: models.FriendRequestID(r.ID.String()),
		Status: models.FriendRequestStatus(r.Status),
		FromUserID: models.UserID(r.FromUserID.String()),
		ToUserID: models.UserID(r.ToUserID.String()),
		CreatedAt: r.CreatedAt,

	}
}

// FromSaveDTO конвертирует dto в FriendRequestRow для INSERT
func FromSaveDTO(d *dto.SaveFriendRequest) FriendRequestRow {
	if d == nil {
		return FriendRequestRow{}
	}

	parsedFromUserID, err := uuid.Parse(string(d.FromUserID))
	if err != nil {
		return FriendRequestRow{}
	}

	parsedToUserID, err := uuid.Parse(string(d.ToUserID))
	if err != nil {
		return FriendRequestRow{}
	}

	return FriendRequestRow{
		FromUserID: parsedFromUserID,
		ToUserID: parsedToUserID,
		Status: uint64(d.Status),
	}
}

// FromUpdateDTO конвертирует доменную модель в FriendRequestRow для UPDATE
func FromUpdateDTO(d *dto.UpdateFriendRequest) FriendRequestRow {
	if d == nil {
		return FriendRequestRow{}
	}

	parsedReqID, err := uuid.Parse(string(d.ReqID))
	if err != nil {
		return FriendRequestRow{}
	}

	return FriendRequestRow{
		ID: parsedReqID,
		Status: uint64(d.Status),
	}
}