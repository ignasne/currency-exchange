package router

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func New(router *mux.Router) *Router {
	r := new(Router)
	r.Router = router
	return r
}
