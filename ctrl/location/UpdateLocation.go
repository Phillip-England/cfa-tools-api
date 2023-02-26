package ctrl

import (
	"context"
	"net/http"
	"time"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateLocation(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		Name   string `json:"name"`
		Number string `json:"number"`
		CSRF   string `json:"_csrf"`
	}

	body := requestBody{}
	lib.GetBody(w, r, &body)

	id := lib.GetURLParam(r.URL.Path)
	locationID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.ResourceNotFound(w)
		return
	}

	err = lib.IsCSRF(body.CSRF)
	if err != nil {
		res.Forbidden(w)
		return
	}

	user := r.Context().Value(model.GetUserKey()).(model.User)

	coll := db.Collection(client, "locations")

	location := model.Location{}
	filter := bson.D{{Key: "_id", Value: locationID}}
	err = coll.FindOne(context.Background(), filter).Decode(&location)
	if err == mongo.ErrNoDocuments {
		res.ResourceNotFound(w)
		return
	} else if err != nil {
		res.ServerError(w, err)
		return
	}

	if location.User != user.ID {
		res.Forbidden(w)
		return
	}

	filter = bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "name", Value: body.Name},
			{Key: "number", Value: body.Number},
			{Key: "updatedAt", Value: time.Now()},
		},
	}}
	_, err = coll.UpdateByID(context.Background(), location.ID, filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)

}
