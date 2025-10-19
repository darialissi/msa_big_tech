package config

import (
	"errors"
)

type jwtEnv struct {
	secret string
}

func JWTConfig() *jwtEnv {
	return &jwtEnv{
		secret: getEnv("JWT_SECRET", ""),
	}
}

func (env *jwtEnv) Validate() error {
	if env.secret == "" {
		return errors.New("No defined JWT secret")
	}
	return nil
}

func (env *jwtEnv) GetSecret() string {
	return env.secret
}
