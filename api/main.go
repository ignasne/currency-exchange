package main

import (
	"github.com/ignasne/currency-exchange/api/api"
	"github.com/ignasne/currency-exchange/api/config"
	"github.com/ignasne/currency-exchange/api/router"
)

func main() {
	config := &config.Main{}
	config.Parse()

	httpAPI := api.New(config.SelfPort)
	_ = router.New(httpAPI.Mux)

	httpAPI.RegisterRoutes()
	httpAPI.Listen()
}
