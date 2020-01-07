package cmd

import (
	"github.com/ignasne/currency-exchange/api/api"
	"github.com/ignasne/currency-exchange/api/config"
	"github.com/ignasne/currency-exchange/api/datasource"
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

	_ = datasource.Connect(cfg.DB)

	httpAPI := api.New(cfg.SelfPort)
	_ = router.New(httpAPI.Mux)

	httpAPI.RegisterRoutes()
	httpAPI.Listen()
}