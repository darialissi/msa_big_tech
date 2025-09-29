package auth

import (
	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases/dto"
)


func (r *Repository) Save(user *dto.SaveCred) (*models.User, error) {
	r.db.Exec(
		"INSERT INTO users (email, password_hash) VALUES (?, ?)",
		user.Email, user.PasswordHash,
	)

	//

	return &models.User{}, nil
}

func (r *Repository) Update(user *dto.UpdateCredSave) (*models.User, error) {
	r.db.Exec(
		"UPDATE users SET nickname=?, bio=?, avatar_url=? WHERE id=?",
		user.Email, user.PasswordHash, user.ID,
	)

	//

	return &models.User{}, nil
}

func (r *Repository) FetchById(userId dto.UserID) (*models.User, error) {
	r.db.Exec(
		"SELECT * FROM users WHERE id=?",
		userId,
	)

	//

	return &models.User{}, nil
}

func (r *Repository) FetchByEmail(email string) (*models.User, error) {
	r.db.Exec(
		"SELECT * FROM users WHERE email=?",
		email,
	)

	//

	return &models.User{}, nil
}