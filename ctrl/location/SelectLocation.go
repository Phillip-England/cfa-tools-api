package ctrl

import (
	"net/http"
	"strings"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SelectLocation(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		res.InvalidRequestMethod(w)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]
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

	filter := bson.D{{Key: "_id", Value: locationID}}
	var location model.LocationResponse
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

	net.HttpCookie(w, "location", id, 1440)

	res.Success(w)

}