package net

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
)

func BadReqeust(w http.ResponseWriter, message string) {
	response := model.SimpleResponse{Message: message}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	w.WriteHeader(400)
	w.Write(jsonBytes)
}