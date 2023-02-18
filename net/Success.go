package net

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
)

func Success(w http.ResponseWriter) {
	response := model.SimpleResponse{Message: "success"}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(200)
	w.Write(jsonBytes)
}
