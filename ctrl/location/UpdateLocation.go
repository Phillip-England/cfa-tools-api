package ctrl

import (
	"net/http"
	"time"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateLocation(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		res.InvalidRequestMethod(w)
		return
	}

	type requestBody struct {
		Name string `json:"name"`
		Number string `json:"number"`
	}

	body := requestBody{}
	net.GetBody(w, r, &body)

	id := net.GetURLParam(r.URL.Path)
	locationID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.ResourceNotFound(w)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "locations")

	location := model.Location{}
	filter := bson.D{{Key: "_id", Value: locationID}}
	err = coll.FindOne(ctx, filter).Decode(&location)
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
		Key:"$set", Value: bson.D{
			{Key: "name", Value: body.Name},
			{Key: "number", Value: body.Number},
			{Key: "updated_at", Value: time.Now()},
		},
	}}
	_, err = coll.UpdateByID(ctx, location.ID, filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)



}