package ctrl

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSingleCares(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	id := lib.GetURLParam(r.URL.Path)
	caresID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.ResourceNotFound(w)
		return
	}


	// location := r.Context().Value(model.GetLocationKey()).(model.Location)

	coll := db.Collection(client, "cares")

	filter := bson.D{{Key: "_id", Value: caresID}}

	var cares model.Cares
	err = coll.FindOne(context.Background(), filter).Decode(&cares)
	if err == mongo.ErrNoDocuments {
		res.ResourceNotFound(w)
		return
	} else if err != nil {
		res.ServerError(w, err)
		return
	}

	httpResponse := model.HttpResponse{
		Message: "success",
		Data:    cares,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}
