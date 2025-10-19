package user

import (
	"context"
	"errors"
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

	pool := r.db.GetQueryEngine(ctx)

	var outRow UserRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) Update(ctx context.Context, in *models.User) (*models.User, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	query := r.sb.
		Update(usersTable).
		Set(usersTableColumnNickname, row.Nickname)

	if row.Bio != "" {
		query = query.Set(usersTableColumnBio, row.Bio)
	}
	if row.AvatarUrl != "" {
		query = query.Set(usersTableColumnAvatarUrl, row.AvatarUrl)
	}

	query = query.
		Where(squirrel.Eq{usersTableColumnID: row.ID}).
		Suffix("RETURNING " + strings.Join(usersTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow UserRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, id models.UserID) (*models.User, error) {

	query := r.sb.
		Select(usersTableColumns...).
		From(usersTable).
		Where(squirrel.Eq{usersTableColumnID: id})

	pool := r.db.GetQueryEngine(ctx)

	var outRow UserRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
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

	pool := r.db.GetQueryEngine(ctx)

	var outRow UserRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
			return nil, nil
		}
		return nil, err
	}

	return ToModel(&outRow), nil
}
