package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
    ErrUniqueViolation = errors.New("unique constraint violation")
)

func IsUniqueViolation(err error) bool {
    var pgErr *pgconn.PgError
    return errors.As(err, &pgErr) && pgErr.Code == "23505"
}