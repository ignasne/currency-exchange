package config

import (
	"github.com/ignasne/currency-exchange/api/logger"
	"os"

	"github.com/caarlos0/env"
)

type Migration struct {
	Path string `env:"APP_MIGRATIONS_PATH" envDefault:"migrations"`
}

func GetMigration() *Migration {
	c := &Migration{}

	err := env.Parse(c)
	if err != nil {
		logger.Get().WithError(err).Fatal("could not parse migration config")
		os.Exit(1)
	}

	return c
}
