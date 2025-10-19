package chat_member

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
)

// ChatMemberRow — «плоская» проекция строки таблицы messages
// Nullable поля представлены sql.Null*.
type ChatMemberRow struct {
	ChatID uuid.UUID `db:"chat_id"`
	UserID uuid.UUID `db:"user_id"`
}

func (row *ChatMemberRow) Values() []any {
	return []any{
		row.ChatID, row.UserID,
	}
}

// ToModel конвертирует ChatMemberRow (sql.Null*) в доменную модель models.ChatMember (*string/*time.Time).
func ToModel(r *ChatMemberRow) *models.ChatMember {
	if r == nil {
		return nil
	}
	return &models.ChatMember{
		ChatID: models.ChatID(r.ChatID.String()),
		UserID: models.UserID(r.UserID.String()),
	}
}

// FromModel конвертирует доменную модель в ChatMemberRow для INSERT/UPDATE/DELETE
func FromModel(m *models.ChatMember) (ChatMemberRow, error) {
	if m == nil {
		return ChatMemberRow{}, fmt.Errorf("model is nil")
	}

	var chatID uuid.UUID
	var userID uuid.UUID
	var err error

	// Универсальная конвертация, поля могут быть пропущены (прописываются явно в запросе репо)
	if string(m.ChatID) != "" {
		chatID, err = uuid.Parse(string(m.ChatID))
		if err != nil {
			return ChatMemberRow{}, fmt.Errorf("invalid chat_id: %w", err)
		}
	}

	if string(m.UserID) != "" {
		userID, err = uuid.Parse(string(m.UserID))
		if err != nil {
			return ChatMemberRow{}, fmt.Errorf("invalid user_id: %w", err)
		}
	}

	return ChatMemberRow{
		ChatID: chatID,
		UserID: userID,
	}, nil
}
