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

func GetAllCares(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	location := r.Context().Value(model.GetLocationKey()).(model.Location)

	coll := db.Collection(client, "cares")

	filter := bson.D{{Key: "location", Value: location.ID}}

	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}
	defer cursor.Close(context.Background())

	var allCares []model.Cares
	for cursor.Next(context.Background()) {
		var cares model.Cares
		if err := cursor.Decode(&cares); err != nil {
			res.ServerError(w, err)
			return
		}
		allCares = append(allCares, cares)
	}

	httpResponse := model.HttpResponse{
		Message: "success",
		Data:    allCares,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}
