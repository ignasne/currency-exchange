package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

func (c *Main) Parse() {
	err := env.Parse(c)
	if err != nil {
		fmt.Printf("Could not parse Config")
		os.Exit(1)
	}
}
