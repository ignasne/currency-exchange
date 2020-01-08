package router

import (
	"errors"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/ignasne/currency-exchange/api/logger"
	"github.com/ignasne/currency-exchange/api/quote"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
)


type quoteRouter struct {}

func (r *Router) RegisterQuoteRoutes(currencies quote.Currencies) {
	router := &quoteRouter{}

	mr := r.Router.PathPrefix("/api").Subrouter()

	mr.HandleFunc("/quote", router.getHandler(currencies)).Methods("GET")
}

func (r *quoteRouter) getHandler(currencies quote.Currencies) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		requestValues, err := getQuoteRequestValues(req.URL.Query())
		if err != nil {
			BadRequest(res, err)
			return
		}

		_, err = validateRequest(requestValues, currencies)
		if err != nil {
			BadRequest(res, err)
			return
		}
	}
}

func validateRequest(requestData *quote.RequestGetStruct, currencies quote.Currencies) (bool, error) {
	if !currencies.Validate(requestData.FromCurrencyCode) || !currencies.Validate(requestData.ToCurrencyCode) {
		// It not recommended to reveal available values in api response as they should be seen in api
		// documentation.
		// But at this time let's show them
		return false, errors.New(fmt.Sprintf("Bad currency value. Available currencies: %s", strings.Join(currencies, ", ")))
	}

	// amount is already validated as integer schema parsing
	if requestData.Amount < 0 {
		// Errors strings can be lower_case_underscore then ui can translate them in anny language they want
		// But it's depends on requirements. So move on with error strings.
		return false, errors.New("Amount should be greater than 0.")
	}

	return true, nil
}

func getQuoteRequestValues(query url.Values) (*quote.RequestGetStruct, error) {
	quoteRequest := &quote.RequestGetStruct{}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err := decoder.Decode(quoteRequest, query)

	if err != nil {
		logger.Get().WithFields(logrus.Fields{"req": query, "error": err}).Info("failed to parse quote url parameters")
		return nil, err
	}

	return quoteRequest, nil
}
