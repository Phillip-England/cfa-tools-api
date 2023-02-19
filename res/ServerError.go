package res

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/phillip-england/go-http/model"
)

func ServerError(w http.ResponseWriter, err error) {


	if os.Getenv("GO_ENV") == "prod" {
		response := model.ErrorResponse{
			Message: "internal server error",
			Error: err.Error(),
		}
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonBytes)
	} else {
		response := model.SimpleResponse{Message: "internal server error"}
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonBytes)
	}

}
