package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		net.MessageResponse(w, "invalid request method", http.StatusBadRequest)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	userResponse := model.UserResponse{
		ID: user.ID,
		Email: user.Email,
	}

	httpResponse := model.HttpResponse{
		Message: "success",
		Data: userResponse,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		net.ServerError(w)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}