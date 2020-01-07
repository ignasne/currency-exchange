package main

import (
	"github.com/ignasne/currency-exchange/api/api"
	"github.com/ignasne/currency-exchange/api/config"
	"github.com/ignasne/currency-exchange/api/datasource"
	"github.com/ignasne/currency-exchange/api/router"
)

func main() {
	cfg := &config.Main{}
	cfg.Parse()

	_ = datasource.Connect(cfg.DB)

	httpAPI := api.New(cfg.SelfPort)
	_ = router.New(httpAPI.Mux)

	httpAPI.RegisterRoutes()
	httpAPI.Listen()
}
