package quote

import (
	"github.com/shopspring/decimal"
)

type apiClientI interface {
	GetRateForCurrencies(fromCurrency string, toCurrency string) (decimal.Decimal, error)
}

type cacheClientI interface {
	Get(key string) *string
	Set(key string, value int, ttl int) bool
}

// Currency service
type Currency struct {
	apiClient   apiClientI
	cacheClient cacheClientI
}

func (c *Currency) GetRate(fromCurrency string, toCurrency string) (*Rate, error) {
	rateDecimal, err := c.apiClient.GetRateForCurrencies(fromCurrency, toCurrency)

	if err != nil {
		return nil, err
	}

	// round rate
	rateNumber, _ := rateDecimal.Round(3).Float64() // round it to nearest

	result := &Rate{value: rateDecimal, roundedValue: rateNumber}

	return result, nil
}

func GetCurrencyService(apiClient apiClientI, cacheClient cacheClientI) *Currency {
	return &Currency{apiClient: apiClient, cacheClient: cacheClient}
}
