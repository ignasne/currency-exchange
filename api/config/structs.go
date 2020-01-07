package config

type Main struct {
	SelfPort int `env:"APP_PORT" envDefault:"3001"`
}
