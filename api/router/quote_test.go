package router

import (
	"fmt"
	gbl "github.com/franela/goblin"
	"github.com/ignasne/currency-exchange/api/quote"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RatesMock struct{}

func (r RatesMock) GetRate(fromCurrency string, toCurrency string) (*quote.Rate, error) {
	result := &quote.Rate{Value: decimal.NewFromFloat32(0.8996851102), RoundedValue: 0.9}
	return result, nil
}

func CheckFailingCase(vars map[string]string) {
	var req *http.Request
	var res *httptest.ResponseRecorder

	ratesMock := RatesMock{}

	currencies := quote.Currencies{
		"USD",
		"EUR",
		"ILS",
	}

	url := "api/quote?"

	for varKey, varValue := range vars {
		url = fmt.Sprintf("%s&%s=%s", url, varKey, varValue)
	}

	req, _ = http.NewRequest(http.MethodGet, url, nil)
	res = httptest.NewRecorder()

	r := quoteRouter{
		RatesService: ratesMock,
	}

	r.getHandler(currencies).ServeHTTP(res, req)

	Expect(res.Code).To(Equal(400))
}

// Tests are provided for api endpoint
// For the production level api there should be tests for repositories, services, clients etc.
func TestQuotesEndpoint(t *testing.T) {
	g := gbl.Goblin(t)

	ratesMock := RatesMock{}

	currencies := quote.Currencies{
		"USD",
		"EUR",
		"ILS",
	}

	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Quote ok request", func() {
		var req *http.Request
		var res *httptest.ResponseRecorder

		g.BeforeEach(func() {
			vars := map[string]string{
				"from_currency_code": "USD",
				"to_currency_code":   "EUR",
				"amount":             "100",
			}

			url := "api/quote?"

			for varKey, varValue := range vars {
				url = fmt.Sprintf("%s&%s=%s", url, varKey, varValue)
			}

			req, _ = http.NewRequest(http.MethodGet, url, nil)
			res = httptest.NewRecorder()
		})

		g.It("should be able to respond success with correct params", func() {
			r := quoteRouter{
				RatesService: ratesMock,
			}

			r.getHandler(currencies).ServeHTTP(res, req)

			bodyBytes, _ := ioutil.ReadAll(res.Body)
			bodyString := string(bodyBytes)

			expected := []byte(`{"data":{"exchange_rate":0.9,"currency_code":"EUR","amount":89}}`)

			Expect(res.Code).To(Equal(200))
			Expect(bodyString).To(Equal(string(expected)))
		})
	})

	g.Describe("Quote fail request", func() {

		g.It("should fail respond if to currency not available", func() {
			vars := map[string]string{
				"from_currency_code": "USD",
				"to_currency_code":   "BTC",
				"amount":             "100",
			}

			CheckFailingCase(vars)
		})

		g.It("should fail respond if from currency not available", func() {
			vars := map[string]string{
				"from_currency_code": "BTC",
				"to_currency_code":   "USD",
				"amount":             "100",
			}

			CheckFailingCase(vars)
		})

		g.It("should fail respond if amount empty", func() {
			vars := map[string]string{
				"from_currency_code": "USD",
				"to_currency_code":   "BTC",
				"amount":             "",
			}

			CheckFailingCase(vars)
		})

		g.It("should fail respond if to currency not given", func() {
			vars := map[string]string{
				"from_currency_code": "USD",
				"amount":             "100",
			}

			CheckFailingCase(vars)
		})

		g.It("should fail respond if from currency not given", func() {
			vars := map[string]string{
				"to_currency_code":   "USD",
				"amount":             "100",
			}

			CheckFailingCase(vars)
		})

		g.It("should fail respond if amount not given", func() {
			vars := map[string]string{
				"from_currency_code": "USD",
				"to_currency_code":   "EUR",
			}

			CheckFailingCase(vars)
		})

		// Amount length should be validated from gui side too
		// Why to convert such a lot money if you do not have :)
		// But if there will be requirement then need to refactor int64 size to bigger
		g.It("should fail respond for a big amount", func() {
			vars := map[string]string{
				"from_currency_code": "USD",
				"to_currency_code":   "BTC",
				"amount":             "10000000000000000000000000000000000000000",
			}

			CheckFailingCase(vars)
		})
	})
}
