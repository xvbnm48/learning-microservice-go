package main

import (
	"net/http"
)

// type jsonResponse struct {
// 	Error   bool   `json:"error"`
// 	Message string `json:"message"`
// 	// if go version is 1.18 using any, if go version is not 1.18 using interface
// 	Data any `json:"data, omitempty"`
// }

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "hit the broker ",
	}

	// out, _ := json.MarshalIndent(payload, "", "\t")
	_ = app.writeJSON(w, http.StatusOK, payload)

}
