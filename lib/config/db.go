package config

import (
	"errors"
	"fmt"
)

type DbEnv struct {
	user     string
	password string
	host     string
	port     string
	db       string
}

func DbConfig(mode string) *DbEnv {
	if mode == "dev" {
		return &DbEnv{
			host:     getEnv("POSTGRES_HOST_DEV", ""),
			port:     getEnv("POSTGRES_PORT_DEV", ""),
			user:     getEnv("POSTGRES_USER_DEV", ""),
			password: getEnv("POSTGRES_PASSWORD_DEV", ""),
			db:       getEnv("POSTGRES_DB_DEV", ""),
		}
	}
	return &DbEnv{
		host:     getEnv("POSTGRES_HOST", ""),
		port:     getEnv("POSTGRES_PORT", ""),
		user:     getEnv("POSTGRES_USER", ""),
		password: getEnv("POSTGRES_PASSWORD", ""),
		db:       getEnv("POSTGRES_DB", ""),
	}
}

func (env *DbEnv) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.user, env.password, env.host, env.port, env.db,
	)
}

func (env *DbEnv) Validate() error {
	if env.host == "" {
		return errors.New("No defined DB host")
	}

	if env.port == "" {
		return errors.New("No defined DB hosportt")
	}

	if env.user == "" {
		return errors.New("No defined DB user")
	}

	if env.password == "" {
		return errors.New("No defined DB password")
	}

	if env.db == "" {
		return errors.New("No defined DB db")
	}

	return nil
}
