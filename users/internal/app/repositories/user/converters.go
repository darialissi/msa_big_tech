package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
)

// UserRow — «плоская» проекция строки таблицы friends
// Nullable поля представлены sql.Null*.
type UserRow struct {
	ID  		uuid.UUID		`db:"id"`
	Nickname    string			`db:"nickname"`
	AvatarUrl   string			`db:"avatar_url"`
	Bio    		string			`db:"bio"`
	CreatedAt 	time.Time       `db:"created_at"`
}

func (row *UserRow) Values() []any {
	return []any{
		row.ID, row.Nickname, row.AvatarUrl, row.Bio, row.CreatedAt,
	}
}

// ToModel конвертирует UserRow (sql.Null*) в доменную модель models.User (*string/*time.Time).
func ToModel(r *UserRow) *models.User {
	if r == nil {
		return nil
	}
	return &models.User{
		ID: models.UserID(r.ID.String()),
		Nickname: r.Nickname,
		AvatarUrl: r.AvatarUrl,
		Bio: r.Bio,
		CreatedAt: r.CreatedAt,
	}
}

// FromModel конвертирует доменную модель в UserRow для INSERT/UPDATE/DELETE
func FromModel(m *models.User) (UserRow, error) {
	if m == nil {
		return UserRow{}, fmt.Errorf("model is nil")
	}
	
	var ID uuid.UUID
	var err error

	// Универсальная конвертация, поля могут быть пропущены (прописываются явно в логике и запросе репо)
	if string(m.ID) != "" {
		ID, err = uuid.Parse(string(m.ID))
		if err != nil {
			return UserRow{}, fmt.Errorf("invalid user_id: %w", err)
		}
	}

	return UserRow{
		ID: ID,
		Nickname: m.Nickname,
		AvatarUrl: m.AvatarUrl,
		Bio: m.Bio,
		CreatedAt: m.CreatedAt,
	}, nil
}