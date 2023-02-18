package ctrl

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetLocation(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		net.MessageResponse(w, "invalid request method", 400)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]
	locationID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "locations")

	filter := bson.D{{Key: "_id", Value: locationID}}
	var location model.LocationResponse
	err = coll.FindOne(ctx, filter).Decode(&location)
	if err != nil {
		net.ServerError(w)
		return
	}

	if location.User != user.ID {
		net.Forbidden(w)
		return
	}

	httpResponse := model.HttpResponse{
		Message: "success",
		Data: location,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		net.ServerError(w)
		return
	}
	
	w.WriteHeader(200)
	w.Write(jsonData)

}