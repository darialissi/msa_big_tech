package auth

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
)

// AuthUserRow — «плоская» проекция строки таблицы auth_users
// Nullable поля представлены sql.Null*.
type AuthUserRow struct {
	ID  			uuid.UUID		`db:"id"`
	Email  			string			`db:"email"`
	PasswordHash  	string			`db:"password_hash"`
	CreatedAt 		time.Time       `db:"created_at"`
}

func (row *AuthUserRow) Values() []any {
	return []any{
		row.ID, row.Email, row.PasswordHash, row.CreatedAt,
	}
}

// ToModel конвертирует AuthUserRow (sql.Null*) в доменную модель models.User (*string/*time.Time).
func ToModel(r *AuthUserRow) *models.User {
	if r == nil {
		return nil
	}
	return &models.User{
		ID: models.UserID(r.ID.String()),
		Email: r.Email,
		PasswordHash: r.PasswordHash,
		CreatedAt: r.CreatedAt,
	}
}

// FromModel конвертирует доменную модель в ChatRow для INSERT/UPDATE/DELETE
func FromModel(m *models.User) (AuthUserRow, error) {
	if m == nil {
		return AuthUserRow{}, fmt.Errorf("model is nil")
	}
	
	var ID uuid.UUID
	var err error

	// Универсальная конвертация, поля могут быть пропущены (прописываются явно в логике и репо)
	if string(m.ID) != "" {
		ID, err = uuid.Parse(string(m.ID))
		if err != nil {
			return AuthUserRow{}, fmt.Errorf("invalid id: %w", err)
		}
	}

	return AuthUserRow{
		ID: ID,
		Email: m.Email,
		PasswordHash: m.PasswordHash,
		CreatedAt: m.CreatedAt,
	}, nil
}