package quote

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type apiClientI interface {
	GetRateForCurrencies(fromCurrency string, toCurrency string) (decimal.Decimal, error)
}

type cacheClientI interface {
	Get(key string) *string
	Set(key string, value string, ttl int) bool
}

// Currency service
type Currency struct {
	apiClient   apiClientI
	cacheClient cacheClientI
}

func (c *Currency) GetRate(fromCurrency string, toCurrency string) (*Rate, error) {
	cacheKey := fmt.Sprintf("%s%s", fromCurrency, toCurrency)
	var rateDecimal decimal.Decimal
	var err error

	// firstly check cache for currency pair
	ratioFromCache := c.cacheClient.Get(cacheKey)

	if ratioFromCache != nil {
		rateDecimal, err = decimal.NewFromString(*ratioFromCache)

		if err != nil {
			return nil, err
		}

		result := c.getResult(rateDecimal)

		return result, nil
	}

	rateDecimal, err = c.apiClient.GetRateForCurrencies(fromCurrency, toCurrency)

	if err != nil {
		return nil, err
	}

	// ttl could be take from config
	// save currency rate to cache
	c.cacheClient.Set(cacheKey, rateDecimal.String(), 10)

	result := c.getResult(rateDecimal)

	return result, nil
}

func (c *Currency) getResult(rateDecimal decimal.Decimal) *Rate {
	// round rate
	rateNumber, _ := rateDecimal.Round(3).Float64()

	result := &Rate{value: rateDecimal, roundedValue: rateNumber}

	return result
}

func GetCurrencyService(apiClient apiClientI, cacheClient cacheClientI) *Currency {
	return &Currency{apiClient: apiClient, cacheClient: cacheClient}
}
