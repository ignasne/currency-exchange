package quote

import (
	"errors"
	"math"
	"time"
)

type apiClientI interface {
	GetRateForCurrencies(fromCurrency string, toCurrency string) (float64, error)
}

type cacheClientI interface {
	Get(key string) string
	Set(key string, value string, ttl time.Time) bool
}

// Currency service
type Currency struct {
	apiClient   apiClientI
	cacheClient cacheClientI
}

func (c *Currency) GetRate(fromCurrency string, toCurrency string) (*Rate, error) {
	rate, err := c.apiClient.GetRateForCurrencies(fromCurrency, toCurrency)

	if err != nil {
		return nil, err
	}

	if rate == -1 {
		return nil, errors.New("currency to convert not found in api response")
	}

	// round rate
	rateNumber := math.Round(rate*1000) / 1000 // round it to nearest

	result := &Rate{value: rate, roundedValue: rateNumber}

	return result, nil
}

func GetCurrencyService(apiClient apiClientI, cacheClient cacheClientI) *Currency {
	currencyService := &Currency{apiClient: apiClient, cacheClient: cacheClient}
	return currencyService
}
