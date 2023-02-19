package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteLocation(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		net.InvalidRequestMethod(w)
		return
	}

	id := net.GetURLParam(r.URL.Path)
	locationID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		net.ResourceNotFound(w)
		return
	}
	
	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "locations")
	
	var location model.Location
	filter := bson.D{{Key: "_id", Value: locationID}}
	err = coll.FindOne(ctx, filter).Decode(&location)
	if err == mongo.ErrNoDocuments {
		net.ResourceNotFound(w)
		return
	} else if err != nil {
		net.ServerError(w)
		return
	}

	if user.ID != location.User {
		net.Forbidden(w)
		return
	}

	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		net.ServerError(w)
		return
	}

	net.Success(w)

}