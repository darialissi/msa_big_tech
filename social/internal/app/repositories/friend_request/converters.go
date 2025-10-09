package friend_request

import (
	"time"
	"fmt"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
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

// FromModel конвертирует доменную модель в FriendRequestRow для INSERT/UPDATE/DELETE
func FromModel(m *models.FriendRequest) (FriendRequestRow, error) {
	if m == nil {
		return FriendRequestRow{}, fmt.Errorf("model is nil")
	}
	
	var ID uuid.UUID
	var fromUserID uuid.UUID
	var toUserID uuid.UUID
	var err error

	// Универсальная конвертация, поля могут быть пропущены (прописываются явно в запросе репо)
	if string(m.ID) != "" {
		ID, err = uuid.Parse(string(m.ID))
		if err != nil {
			return FriendRequestRow{}, fmt.Errorf("invalid id: %w", err)
		}
	}

	if string(m.FromUserID) != "" {
		fromUserID, err = uuid.Parse(string(m.FromUserID))
		if err != nil {
			return FriendRequestRow{}, fmt.Errorf("invalid from_user_id: %w", err)
		}
	}

	if string(m.ToUserID) != "" {
		toUserID, err = uuid.Parse(string(m.ToUserID))
		if err != nil {
			return FriendRequestRow{}, fmt.Errorf("invalid to_user_id: %w", err)
		}
	}

	return FriendRequestRow{
		ID: ID,
		FromUserID: fromUserID,
		ToUserID: toUserID,
		Status: uint64(m.Status),
	}, nil
}