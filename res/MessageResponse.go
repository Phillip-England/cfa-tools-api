package res

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
)

func MessageResponse(w http.ResponseWriter, message string, status int) (err error) {
	response := model.SimpleResponse{Message: message}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	w.Write(jsonBytes)
	return nil
}