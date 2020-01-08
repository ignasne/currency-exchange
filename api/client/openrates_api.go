package client

import (
	"encoding/json"
	"errors"
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

func (c *OpenRatesAPIClient) GetRateForCurrencies(fromCurrency string, toCurrency string) (float64, error) {
	rel := &url.URL{Path: "/latest"}
	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return -1, errors.New("fail create request to open rates api")
	}

	q := url.Values{}
	q.Add("base", fromCurrency)
	req.URL.RawQuery = q.Encode()

	res, err := c.Client.Do(req)

	if err != nil {
		return -1, err
	}

	defer res.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return -1, err
	}

	if len(result) == 0 {
		return -1, errors.New("empty result section in open rates fetch")
	}

	ratesList := result["rates"].(map[string]interface{})

	for key, value := range ratesList {
		if key == toCurrency {
			return value.(float64), nil
		}
	}

	return -1, nil
}
