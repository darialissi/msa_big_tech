package message

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
)

// MessageRow — «плоская» проекция строки таблицы messages
// Nullable поля представлены sql.Null*.
type MessageRow struct {
	ID        uuid.UUID `db:"id"`
	ChatID    uuid.UUID `db:"chat_id"`
	SenderID  uuid.UUID `db:"sender_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func (row *MessageRow) Values() []any {
	return []any{
		row.ID, row.ChatID, row.SenderID, row.Text, row.CreatedAt,
	}
}

// ToModel конвертирует MessageRow (sql.Null*) в доменную модель models.Message (*string/*time.Time).
func ToModel(r *MessageRow) *models.Message {
	if r == nil {
		return nil
	}
	return &models.Message{
		ID:        models.MessageID(r.ID.String()),
		ChatID:    models.ChatID(r.ChatID.String()),
		SenderID:  models.UserID(r.SenderID.String()),
		Text:      r.Text,
		CreatedAt: r.CreatedAt,
	}
}

// FromModel конвертирует доменную модель в MessageRow для INSERT/UPDATE/DELETE
func FromModel(m *models.Message) (MessageRow, error) {
	if m == nil {
		return MessageRow{}, fmt.Errorf("model is nil")
	}

	var ID uuid.UUID
	var chatID uuid.UUID
	var senderID uuid.UUID
	var err error

	// Универсальная конвертация, поля могут быть пропущены (прописываются явно в запросе репо)
	if string(m.ID) != "" {
		ID, err = uuid.Parse(string(m.ID))
		if err != nil {
			return MessageRow{}, fmt.Errorf("invalid id: %w", err)
		}
	}

	if string(m.ChatID) != "" {
		chatID, err = uuid.Parse(string(m.ChatID))
		if err != nil {
			return MessageRow{}, fmt.Errorf("invalid chat_id: %w", err)
		}
	}

	if string(m.SenderID) != "" {
		senderID, err = uuid.Parse(string(m.SenderID))
		if err != nil {
			return MessageRow{}, fmt.Errorf("invalid sender_id: %w", err)
		}
	}

	return MessageRow{
		ID:       ID,
		ChatID:   chatID,
		SenderID: senderID,
		Text:     m.Text,
	}, nil
}
