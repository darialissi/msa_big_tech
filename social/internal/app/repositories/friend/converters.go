package friend

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
)

// FriendRow — «плоская» проекция строки таблицы friends
// Nullable поля представлены sql.Null*.
type FriendRow struct {
	UserID    uuid.UUID `db:"user_id"`
	FriendID  uuid.UUID `db:"friend_id"`
	CreatedAt time.Time `db:"created_at"`
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
		UserID:    models.UserID(r.UserID.String()),
		FriendID:  models.UserID(r.FriendID.String()),
		CreatedAt: r.CreatedAt,
	}
}

// FromModel конвертирует доменную модель в FriendRow для INSERT/UPDATE/DELETE
func FromModel(m *models.UserFriend) (FriendRow, error) {
	if m == nil {
		return FriendRow{}, fmt.Errorf("model is nil")
	}

	var userID uuid.UUID
	var friendID uuid.UUID
	var err error

	// Универсальная конвертация, поля могут быть пропущены (прописываются явно в запросе репо)
	if string(m.UserID) != "" {
		userID, err = uuid.Parse(string(m.UserID))
		if err != nil {
			return FriendRow{}, fmt.Errorf("invalid user_id: %w", err)
		}
	}

	if string(m.FriendID) != "" {
		friendID, err = uuid.Parse(string(m.FriendID))
		if err != nil {
			return FriendRow{}, fmt.Errorf("invalid friend_id: %w", err)
		}
	}

	return FriendRow{
		UserID:   userID,
		FriendID: friendID,
	}, nil
}
