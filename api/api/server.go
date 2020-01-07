package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Api struct {
	HTTPPort int
	Mux      *mux.Router
}

func New(port int) *Api {
	srv := new(Api)
	srv.HTTPPort = port
	srv.Mux = mux.NewRouter()
	srv.Mux.StrictSlash(true)

	return srv
}

func (a *Api) RegisterRoutes() {
	a.Mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.NotFound(res, req)
	}).Methods("GET")
	a.Mux.Handle("/self-test", &SelfTestAction{Version: "1"}).Methods("GET")
}

func (a *Api) Listen() {
	handler := cors.AllowAll().Handler(a.Mux)
	err := http.ListenAndServe(fmt.Sprintf(":%d", a.HTTPPort), handler)

	if err != nil {
		fmt.Println("Could not initialize HTTP server")

		os.Exit(1)
	}
}

func Ok(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(res).Encode(data)

	if err != nil {
		fmt.Println("Could not encode json data")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
