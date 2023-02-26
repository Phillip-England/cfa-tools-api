package ctrl

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateLocation(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		Name   string `json:"name"`
		Number string `json:"number"`
		CSRF   string `json:"_csrf"`
	}

	body := requestBody{}
	err := lib.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	err = lib.IsCSRF(body.CSRF)
	if err != nil {
		res.Forbidden(w)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	location, err := model.BuildLocation(user.ID, body.Name, body.Number)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	coll := db.Collection(client, "locations")

	_, err = coll.InsertOne(context.Background(), location)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)
}
