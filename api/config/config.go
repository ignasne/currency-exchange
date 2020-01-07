package config

import (
	"github.com/ignasne/currency-exchange/api/logger"
	"os"

	"github.com/caarlos0/env"
)

func (c *Main) Parse() {
	err := env.Parse(c)
	if err != nil {
		logger.Get().WithError(err).Fatal("Could not parse Config")
		os.Exit(1)
	}
}
