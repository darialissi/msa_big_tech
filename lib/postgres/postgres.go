package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/darialissi/msa_big_tech/lib/config"
)


func NewPostgresConnection(ctx context.Context, cf *config.DbEnv) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(cf.DSN())
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect: %w", err)
	}

	return pool, nil
}