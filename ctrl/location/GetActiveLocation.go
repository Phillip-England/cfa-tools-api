package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetActiveLocation(w http.ResponseWriter, r *http.Request, db model.Db) {

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	cookie, err := r.Cookie("location")
	if err != nil {
		res.BadReqeust(w, "no active location")
		return
	}

	locationID, err := primitive.ObjectIDFromHex(cookie.Value)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	coll := db.Collection("locations")

	filter := bson.D{{Key: "_id", Value: locationID}}
	var location model.LocationResponse
	err = coll.FindOne(db.Ctx, filter).Decode(&location)
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

	httpResponse := model.HttpResponse{
		Message: "success",
		Data:    location,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}
