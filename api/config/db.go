package config

import (
	"github.com/ignasne/currency-exchange/api/logger"
	"os"

	"github.com/caarlos0/env"
)

type DB struct {
	Host string `env:"DB_HOST,required"`
	Port int    `env:"DB_PORT" envDefault:"3306"`
	User string `env:"DB_USERNAME" envDefault:"root"`
	Pass string `env:"DB_PASSWORD" envDefault:""`
	Name string `env:"DB_DATABASE,required"`
}

func GetDB() *DB {
	c := &DB{}
	err := env.Parse(c)
	if err != nil {
		logger.Get().WithError(err).Fatal("Could not parse db config")
		os.Exit(1)
	}

	return c
}
