package res

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
)

func Forbidden(w http.ResponseWriter) {
	response := model.SimpleResponse{Message: "forbidden resource"}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(401)
	w.Write(jsonBytes)
}
