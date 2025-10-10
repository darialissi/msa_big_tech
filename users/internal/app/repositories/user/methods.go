package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
)


func (r *Repository) Save(ctx context.Context, in *models.User) (*models.User, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	if v1, v2 := row.ID.String(), row.Nickname; v1 == "" || v2 == "" {
		return nil, fmt.Errorf(
			"invalid args: row.ID=%s, row.Nickname=%s", 
			v1,
			v2,
		)
	}

	query := r.sb.
		Insert(usersTable).
		Columns(
			usersTableColumnID, 
			usersTableColumnNickname,
			usersTableColumnAvatarUrl, 
			usersTableColumnBio,
			).
		Values(row.ID, row.Nickname, row.AvatarUrl, row.Bio).
		Suffix("RETURNING " + strings.Join(usersTableColumns, ","))

	var outRow UserRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) Update(ctx context.Context, in *models.User) (*models.User, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	if v1, v2 := row.ID.String(), row.Nickname; v1 == "" || v2 == "" {
		return nil, fmt.Errorf(
			"invalid args: row.ID=%s, row.Nickname=%s", 
			v1,
			v2,
		)
	}

	query := r.sb.
		Update(usersTable).
		Set(usersTableColumnNickname, row.Nickname).
		Set(usersTableColumnAvatarUrl, row.AvatarUrl).
		Set(usersTableColumnBio, row.Bio).
		Where(squirrel.Eq{usersTableColumnID: row.ID}).
		Suffix("RETURNING " + strings.Join(usersTableColumns, ","))

	var outRow UserRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, id models.UserID) (*models.User, error) {

	query := r.sb.
		Select(usersTableColumns...).
		From(usersTable).
		Where(squirrel.Eq{usersTableColumnID: id})

	var outRow UserRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // Запись не найдена
			return nil, nil
		}
    	return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchByNickname(ctx context.Context, nickname string) (*models.User, error) {

	query := r.sb.
		Select(usersTableColumns...).
		From(usersTable).
		Where(squirrel.Eq{usersTableColumnNickname: nickname})

	var outRow UserRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // Запись не найдена
			return nil, nil
		}
    	return nil, err
	}

	return ToModel(&outRow), nil
}