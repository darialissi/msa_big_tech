package config

import (
	"errors"
)


type appEnv struct {
	mode string
}

func AppConfig() *appEnv {
	return &appEnv{
		mode: getEnv("APP_MODE", ""),
	}
}

func (env *appEnv) Validate() error {
	if env.mode == "" {
		return errors.New("No defined APP mode")
	}
	return nil
}

func (env *appEnv) GetMode() string {
	return env.mode
}