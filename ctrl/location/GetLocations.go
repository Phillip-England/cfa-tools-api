package ctrl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"go.mongodb.org/mongo-driver/bson"
)

func GetLocations(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		net.MessageResponse(w, "invalid request method", 400)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "locations")

	filter := bson.D{{Key: "user", Value: user.ID}}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		net.ServerError(w)
		return
	}
	defer cursor.Close(ctx)

	var locations []model.LocationResponse
	for cursor.Next(ctx) {
		log.Println("hit")
		var location model.LocationResponse
		if err := cursor.Decode(&location); err != nil {
			net.ServerError(w)
			return
		}
		locations = append(locations, location)
	}

	httpResponse := model.HttpResponse{
		Message: "success",
		Data: locations,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		net.ServerError(w)
		return
	}
	
	w.WriteHeader(200)
	w.Write(jsonData)

}