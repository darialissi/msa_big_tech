package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/darialissi/msa_big_tech/auth/internal/app/models"
)

func (r *Repository) Save(ctx context.Context, in *models.User) (*models.User, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	query := r.sb.
		Insert(authUsersTable).
		Columns(
			authUsersTableColumnEmail,
			authUsersTableColumnPasswordHash,
		).
		Values(row.Email, row.PasswordHash).
		Suffix("RETURNING " + strings.Join(authUsersTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow AuthUserRow
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
		Update(authUsersTable).
		Set(authUsersTableColumnEmail, row.Email).
		Set(authUsersTableColumnPasswordHash, row.PasswordHash).
		Where(squirrel.Eq{authUsersTableColumnID: row.ID}).
		Suffix("RETURNING " + strings.Join(authUsersTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow AuthUserRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, userId models.UserID) (*models.User, error) {

	query := r.sb.
		Select(authUsersTableColumns...).
		From(authUsersTable).
		Where(squirrel.Eq{authUsersTableColumnID: string(userId)})

	pool := r.db.GetQueryEngine(ctx)

	var outRow AuthUserRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
			return nil, nil
		}
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchByEmail(ctx context.Context, email string) (*models.User, error) {

	query := r.sb.
		Select(authUsersTableColumns...).
		From(authUsersTable).
		Where(squirrel.Eq{authUsersTableColumnEmail: email})

	pool := r.db.GetQueryEngine(ctx)

	var outRow AuthUserRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
			return nil, nil
		}
		return nil, err
	}

	return ToModel(&outRow), nil
}
