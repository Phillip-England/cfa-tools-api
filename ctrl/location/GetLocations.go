package ctrl

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLocations(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(model.GetUserKey()).(model.User)

	coll := db.Collection(client, "locations")

	filter := bson.D{{Key: "user", Value: user.ID}}
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}
	defer cursor.Close(context.Background())

	var locations []model.LocationResponse
	for cursor.Next(context.Background()) {
		var location model.LocationResponse
		if err := cursor.Decode(&location); err != nil {
			res.ServerError(w, err)
			return
		}
		locations = append(locations, location)
	}

	httpResponse := model.HttpResponse{
		Message: "success",
		Data:    locations,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}
