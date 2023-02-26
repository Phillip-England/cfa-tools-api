package ctrl

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteLocation(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	id := lib.GetURLParam(r.URL.Path)
	locationID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.ResourceNotFound(w)
		return
	}

	user := r.Context().Value(model.GetUserKey()).(model.User)

	coll := db.Collection(client, "locations")

	var location model.Location
	filter := bson.D{{Key: "_id", Value: locationID}}
	err = coll.FindOne(context.Background(), filter).Decode(&location)
	if err == mongo.ErrNoDocuments {
		res.ResourceNotFound(w)
		return
	} else if err != nil {
		res.ServerError(w, err)
		return
	}

	if user.ID != location.User {
		res.Forbidden(w)
		return
	}

	_, err = coll.DeleteOne(context.Background(), filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)

}
