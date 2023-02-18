package net

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
)

func MessageResponse(w http.ResponseWriter, message string, status int) {
	response := model.SimpleResponse{Message: message}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		ServerError(w)
	}
	w.WriteHeader(status)
	w.Write(jsonBytes)
}