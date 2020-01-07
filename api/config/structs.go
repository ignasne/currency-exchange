package config

type Main struct {
	SelfPort int `env:"APP_PORT" envDefault:"8080"`
	DB       *DB
}
