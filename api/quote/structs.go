package quote

type RequestGetStruct struct {
	FromCurrencyCode string `schema:"from_currency_code,required"`
	ToCurrencyCode   string `schema:"to_currency_code,required"`
	Amount           int    `schema:"amount,required"`
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
