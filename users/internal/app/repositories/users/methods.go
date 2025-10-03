package users

import (
	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases/dto"
)


func (r *Repository) Save(user *dto.CreateUser) (*models.User, error) {
	r.db.Exec(
		"INSERT INTO users (nickname, bio, avatar_url) VALUES (?, ?, ?)",
		user.Nickname, user.Bio, user.Avatar,
	)

	//

	return &models.User{}, nil
}

func (r *Repository) Update(user *dto.UpdateUser) (*models.User, error) {
	r.db.Exec(
		"UPDATE users SET nickname=?, bio=?, avatar_url=? WHERE id=?",
		user.Nickname, user.Bio, user.Avatar, user.ID,
	)

	//

	return &models.User{}, nil
}

func (r *Repository) FetchById(id dto.UserID) (*models.User, error) {
	r.db.Exec(
		"SELECT * FROM users WHERE id=?",
		id,
	)

	//

	return &models.User{}, nil
}

func (r *Repository) FetchByNickname(nickname string) (*models.User, error) {
	r.db.Exec(
		"SELECT * FROM users WHERE nickname=?",
		nickname,
	)

	//

	return &models.User{}, nil
}