package ctrl

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	user := r.Context().Value(model.GetUserKey()).(model.User)

	location, err := model.BuildLocation(user.ID, body.Name, body.Number)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	coll := db.Collection(client, "locations")

	result, err := coll.InsertOne(context.Background(), location)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	id := result.InsertedID.(primitive.ObjectID)
	location.ID = id
	

	httpResponse := model.HttpResponse{
		Message: "success",
		Data:    location,
		CSRF:    nil,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}
