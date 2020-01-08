package config

import (
	"github.com/ignasne/currency-exchange/api/quote"
)

func GetCurrencies() quote.Currencies {
	// Currencies can be served from env variables to by passed on deploy from infrastructure
	// or served from database
	c := quote.Currencies{
		"USD",
		"EUR",
		"ILS",
	}

	return c
}
