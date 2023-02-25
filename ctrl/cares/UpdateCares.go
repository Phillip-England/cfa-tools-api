package ctrl

import (
	"net/http"
	"time"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateCares(w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		LocationID        string `json:"locationID"`
		GuestName         string `json:"guestName"`
		OrderNumber       string `json:"orderNumber"`
		Incident          string `json:"incident"`
		ReplacementAction string `json:"replacementAction"`
		CSRF              string `json:"_csrf"`
	}

	id := net.GetURLParam(r.URL.Path)
	caresID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.ResourceNotFound(w)
		return
	}

	body := requestBody{}
	err = net.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	err = net.IsCSRF(body.CSRF)
	if err != nil {
		res.Forbidden(w)
		return
	}

	const locationKey model.ContextKey = "location"
	location := r.Context().Value(locationKey).(model.Location)

	ctx, client, disconnect := db.Connect()
	coll := db.Collection(client, "cares")
	defer disconnect()

	cares := model.Cares{}
	filter := bson.D{{Key: "_id", Value: caresID}}
	err = coll.FindOne(ctx, filter).Decode(&cares)
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

	filter = bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "guestName", Value: body.GuestName},
			{Key: "orderNumber", Value: body.OrderNumber},
			{Key: "incident", Value: body.Incident},
			{Key: "replacementAction", Value: body.ReplacementAction},
			{Key: "updatedAt", Value: time.Now()},
		},
	}}
	_, err = coll.UpdateByID(ctx, caresID, filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}
	
	res.Success(w)

}
