package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteLocation(w http.ResponseWriter, r *http.Request, db model.Db) {

	id := net.GetURLParam(r.URL.Path)
	locationID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.ResourceNotFound(w)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	coll := db.Collection("locations")

	var location model.Location
	filter := bson.D{{Key: "_id", Value: locationID}}
	err = coll.FindOne(db.Ctx, filter).Decode(&location)
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

	_, err = coll.DeleteOne(db.Ctx, filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)

}
