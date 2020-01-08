package cmd

import (
	"github.com/ignasne/currency-exchange/api/api"
	"github.com/ignasne/currency-exchange/api/client"
	"github.com/ignasne/currency-exchange/api/config"
	"github.com/ignasne/currency-exchange/api/datasource"
	"github.com/ignasne/currency-exchange/api/logger"
	"github.com/ignasne/currency-exchange/api/quote"
	"github.com/ignasne/currency-exchange/api/router"
	"github.com/spf13/cobra"
)

var httpServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		runHttpServer()
	},
}

func init() {
	rootCmd.AddCommand(httpServerCmd)
}

func runHttpServer() {
	cfg := &config.Main{}
	cfg.Parse()

	db := datasource.Connect(cfg.DB)

	httpAPI := api.New(cfg.SelfPort)
	r := router.New(httpAPI.Mux)

	openRatesApiClient, err := client.NewOpenRatesAPIClient(cfg.RatesAPIURL)

	if err != nil {
		logger.Get().WithError(err).Fatal("failed to initialize open rates api client")
	}

	ratesService := quote.GetCurrencyService(openRatesApiClient, &datasource.CacheDB{DB: db})

	r.RegisterQuoteRoutes(cfg.Currencies, ratesService)

	httpAPI.RegisterRoutes()
	httpAPI.Listen()
}
