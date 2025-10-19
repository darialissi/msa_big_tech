package chat

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
)

// ChatRow — «плоская» проекция строки таблицы messages
// Nullable поля представлены sql.Null*.
type ChatRow struct {
	ID        uuid.UUID `db:"id"`
	CreatorID uuid.UUID `db:"creator_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (row *ChatRow) Values() []any {
	return []any{
		row.ID, row.CreatorID, row.CreatedAt,
	}
}

// ToModel конвертирует ChatRow (sql.Null*) в доменную модель models.Message (*string/*time.Time).
func ToModel(r *ChatRow) *models.DirectChat {
	if r == nil {
		return nil
	}
	return &models.DirectChat{
		ID:        models.ChatID(r.ID.String()),
		CreatorID: models.UserID(r.CreatorID.String()),
		CreatedAt: r.CreatedAt,
	}
}

// FromModel конвертирует доменную модель в ChatRow для INSERT/UPDATE/DELETE
func FromModel(m *models.DirectChat) (ChatRow, error) {
	if m == nil {
		return ChatRow{}, fmt.Errorf("model is nil")
	}

	var ID uuid.UUID
	var creatorID uuid.UUID
	var err error

	// Универсальная конвертация, поля могут быть пропущены (прописываются явно в запросе репо)
	if string(m.ID) != "" {
		ID, err = uuid.Parse(string(m.ID))
		if err != nil {
			return ChatRow{}, fmt.Errorf("invalid id: %w", err)
		}
	}

	if string(m.CreatorID) != "" {
		creatorID, err = uuid.Parse(string(m.CreatorID))
		if err != nil {
			return ChatRow{}, fmt.Errorf("invalid chat_id: %w", err)
		}
	}

	return ChatRow{
		ID:        ID,
		CreatorID: creatorID,
	}, nil
}
