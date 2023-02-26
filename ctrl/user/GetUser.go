package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	userResponse := model.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	token, err := lib.GetCSRF()
	if err != nil {
		res.ServerError(w, err)
		return
	}

	httpResponse := model.HttpResponse{
		Message: "success",
		Data:    userResponse,
		CSRF:    token,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}
