package auth

import (
	"context"
	"errors"
	"fmt"
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

	if v1, v2 := row.Email, row.PasswordHash; v1 == "" || v2 == "" {
		return nil, fmt.Errorf(
			"invalid args: row.Email=%s row.PasswordHash=%s", 
			v1,
			v2,
		)
	}

	query := r.sb.
		Insert(authUsersTable).
		Columns(
			authUsersTableColumnEmail, 
			authUsersTableColumnPasswordHash,
			).
		Values(row.Email, row.PasswordHash).
		Suffix("RETURNING " + strings.Join(authUsersTableColumns, ","))

	var outRow AuthUserRow
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

	if v1, v2, v3 := row.ID.String(), row.Email, row.PasswordHash; v1 == "" || v2 == "" || v3 == "" {
		return nil, fmt.Errorf(
			"invalid args: row.ID=%s row.Email=%s row.PasswordHash=%s", 
			v1,
			v2,
			v3,
		)
	}

	query := r.sb.
		Update(authUsersTable).
		Set(authUsersTableColumnEmail, row.Email).
		Set(authUsersTableColumnPasswordHash, row.PasswordHash).
		Where(squirrel.Eq{authUsersTableColumnID: row.ID}).
		Suffix("RETURNING " + strings.Join(authUsersTableColumns, ","))

	var outRow AuthUserRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, userId models.UserID) (*models.User, error) {

	query := r.sb.
		Select(authUsersTableColumns...).
		From(authUsersTable).
		Where(squirrel.Eq{authUsersTableColumnID: string(userId)})

	var outRow AuthUserRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
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

	var outRow AuthUserRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
			return nil, nil
		}
    	return nil, err
	}

	return ToModel(&outRow), nil
}