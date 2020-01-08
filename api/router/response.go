package router

import "net/http"
import resp "github.com/nicklaw5/go-respond"

type ResourceResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}


// BadRequest writes 402 status code with error message to http response
func BadRequest(r http.ResponseWriter, err error) {
	resp.NewResponse(r).BadRequest(&ResourceResponse{Error: err.Error()})
}

// Ok writes 200 response with resource data
func Ok(r http.ResponseWriter, v interface{}) {
	resp.NewResponse(r).Ok(&ResourceResponse{Data: v})
}

// ServerError writes 501 status code with error message
func ServerError(r http.ResponseWriter, err error) {
	resp.NewResponse(r).InternalServerError(&ResourceResponse{Error: err.Error()})
}
