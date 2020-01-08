package quote

type RequestGetStruct struct {
	FromCurrencyCode string `schema:"from_currency_code,required"`
	ToCurrencyCode   string `schema:"to_currency_code,required"`
	Amount           int64    `schema:"amount,required"`
}

type ResponseGet struct {
	ExchangeRate float64 `json:"exchange_rate"`
	CurrencyCode string  `json:"currency_code"`
	Amount       int64   `json:"amount"`
}

type Rate struct {
	value float64
	roundedValue float64
}

type Currencies []string

func (c Currencies) Validate(currency string) bool {
	for i := range c {
		if c[i] == currency {
			return true
		}
	}

	return false
}
