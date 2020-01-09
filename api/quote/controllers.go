package quote

import "github.com/shopspring/decimal"

type quotesFinderI interface {
	GetRate(fromCurrency string, toCurrency string) (*Rate, error)
}

type Controller struct {
	Rates quotesFinderI
}

func (c *Controller) Get(fromCurrencyCode string, toCurrencyCode string, amount int64) (*ResponseGet, error) {

	rate, err := c.Rates.GetRate(fromCurrencyCode, toCurrencyCode)

	if err != nil || rate == nil {
		return nil, err
	}

	amountDecimal := decimal.NewFromInt(amount)

	// convert Value to cents
	rateResult, _ := rate.Value.Mul(amountDecimal).Float64()

	response := &ResponseGet{
		ExchangeRate: rate.RoundedValue,
		CurrencyCode: toCurrencyCode,
		Amount: int64(rateResult),
	}

	return response, nil
}
