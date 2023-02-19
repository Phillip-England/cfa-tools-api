package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		res.InvalidRequestMethod(w)
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
		res.ServerError(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}