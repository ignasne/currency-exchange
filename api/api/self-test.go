package api

import (
	"encoding/json"
	"net/http"
)

type SelfTestAction struct {
	Version string
}

func (s *SelfTestAction) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(res).Encode(s)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
