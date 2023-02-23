package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCares(w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		LocationID        string `json:"locationID"`
		GuestName         string `json:"guestName"`
		OrderNumber       string `json:"orderNumber"`
		Incident          string `json:"incident"`
		ReplacementAction string `json:"replacementAction"`
		CSRF              string `json:"_csrf"`
	}

	body := requestBody{}
	err := net.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	err = net.IsCSRF(body.CSRF)
	if err != nil {
		res.Forbidden(w)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	locationID, err := primitive.ObjectIDFromHex(body.LocationID)
	if err != nil {
		res.ServerError(w, err)
		return
	}

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

	cares := model.Cares{
		User:              user.ID,
		Location:          locationID,
		GuestName:         body.GuestName,
		OrderNumber:       body.OrderNumber,
		Incident:          body.Incident,
		ReplacementAction: body.ReplacementAction,
	}
	cares.Timestamp()

	coll = db.Collection(client, "cares")
	err = cares.SetReplacementCode(ctx, coll)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	_, err = coll.InsertOne(ctx, cares)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)

}
