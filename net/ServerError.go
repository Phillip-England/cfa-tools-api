package net

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
)

func ServerError(w http.ResponseWriter) {
	response := model.SimpleResponse{Message: "internal server error"}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonBytes)
}
