package net

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
)

func Unauthorized(w http.ResponseWriter) {
	response := model.SimpleResponse{Message: "unauthorized"}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(401)
	w.Write(jsonBytes)
}