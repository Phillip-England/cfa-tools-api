package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
)

func GetLocations(w http.ResponseWriter, r *http.Request, db model.Db) {

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	coll := db.Collection("locations")

	filter := bson.D{{Key: "user", Value: user.ID}}
	cursor, err := coll.Find(db.Ctx, filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}
	defer cursor.Close(db.Ctx)

	var locations []model.LocationResponse
	for cursor.Next(db.Ctx) {
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
