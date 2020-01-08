package quote

type quotesFinderI interface {
	GetRate(fromCurrency string, toCurrency string) (*Rate, error)
}

//type cacheServiceI interface {
//	Get(key string)
//	Set(key string, value string, ttl civil.DateTime)
//}

type Controller struct {
	Rates quotesFinderI
	//CacheService cacheServiceI
}

func (c *Controller) Get(fromCurrencyCode string, toCurrencyCode string, amount int64) (*ResponseGet, error) {

	rate, err := c.Rates.GetRate(fromCurrencyCode, toCurrencyCode)

	if err != nil || rate == nil {
		return nil, err
	}

	// convert value to cents
	rateResult := int64(rate.value * float64(amount))

	response := &ResponseGet{
		ExchangeRate: rate.roundedValue,
		CurrencyCode: toCurrencyCode,
		Amount: rateResult,
	}

	return response, nil
}
