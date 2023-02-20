package res

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
)

func ResourceNotFound(w http.ResponseWriter) {
	response := model.SimpleResponse{Message: "resource not found"}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(404)
	w.Write(jsonBytes)
}
