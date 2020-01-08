package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"net/http"
	"net/url"
)

type OpenRatesAPIClient struct {
	BaseURL *url.URL
	Client  *http.Client
}

func NewOpenRatesAPIClient(apiURL string) (*OpenRatesAPIClient, error) {
	u, err := url.Parse(apiURL)

	if err != nil {
		return nil, errors.New("failed to parse open rates api URL")
	}

	return &OpenRatesAPIClient{
		BaseURL: u,
		Client:  &http.Client{},
	}, nil
}

func (c *OpenRatesAPIClient) GetRateForCurrencies(fromCurrency string, toCurrency string) (decimal.Decimal, error) {
	rel := &url.URL{Path: "/latest"}
	u := c.BaseURL.ResolveReference(rel)
	var currencyRate decimal.Decimal

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return currencyRate, errors.New("fail create request to open rates api")
	}

	q := url.Values{}
	q.Add("base", fromCurrency)
	req.URL.RawQuery = q.Encode()

	res, err := c.Client.Do(req)

	if err != nil {
		return currencyRate, err
	}

	defer res.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return currencyRate, err
	}

	if len(result) == 0 {
		return currencyRate, errors.New("empty result section in open rates fetch")
	}

	ratesList := result["rates"].(map[string]interface{})
	currencyRateText := ""

	for key, value := range ratesList {
		if key == toCurrency {
			currencyRateText = fmt.Sprintf("%v", value)
		}
	}

	if len(currencyRateText) == 0 {
		return currencyRate, errors.New(fmt.Sprintf("rate %s not found for %s currency", toCurrency, fromCurrency))
	}

	currencyRate, err = decimal.NewFromString(currencyRateText)

	if err != nil {
		return currencyRate, errors.New(fmt.Sprintf("could not convert currency %s rate", toCurrency))
	}

	return currencyRate, nil
}
