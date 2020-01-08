package config

import "github.com/ignasne/currency-exchange/api/quote"

type Main struct {
	SelfPort int `env:"APP_PORT" envDefault:"8080"`
	DB       *DB
	Currencies quote.Currencies
}
