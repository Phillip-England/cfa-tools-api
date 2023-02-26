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

func DeleteCares(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	id := lib.GetURLParam(r.URL.Path)
	caresID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.ResourceNotFound(w)
		return
	}

	const locationKey model.ContextKey = "location"
	location := r.Context().Value(locationKey).(model.Location)

	coll := db.Collection(client, "cares")

	var cares model.Cares
	filter := bson.D{{Key: "_id", Value: caresID}}
	err = coll.FindOne(context.Background(), filter).Decode(&cares)
	if err == mongo.ErrNoDocuments {
		res.ResourceNotFound(w)
		return
	} else if err != nil {
		res.ServerError(w, err)
		return
	}

	if cares.Location != location.ID {
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
