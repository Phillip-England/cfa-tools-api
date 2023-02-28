package ctrl

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCares(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		LocationID        string `json:"locationID"`
		GuestName         string `json:"guestName"`
		OrderNumber       string `json:"orderNumber"`
		Incident          string `json:"incident"`
		ReplacementAction string `json:"replacementAction"`
		CSRF              string `json:"_csrf"`
	}

	body := requestBody{}
	err := lib.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	err = lib.IsCSRF(body.CSRF)
	if err != nil {
		res.Forbidden(w)
		return
	}

	user := r.Context().Value(model.GetUserKey()).(model.User)
	location := r.Context().Value(model.GetLocationKey()).(model.Location)

	cares := model.Cares{
		User:              user.ID,
		Location:          location.ID,
		GuestName:         body.GuestName,
		OrderNumber:       body.OrderNumber,
		Incident:          body.Incident,
		ReplacementAction: body.ReplacementAction,
	}
	cares.Timestamp()

	coll := db.Collection(client, "cares")
	err = cares.SetReplacementCode(context.Background(), coll)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	result, err := coll.InsertOne(context.Background(), cares)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	cares.ID = result.InsertedID.(primitive.ObjectID)

	httpResponse := model.HttpResponse{
		Message: "success",
		Data:    cares,
		CSRF:    nil,
	}

	jsonData, err := json.Marshal(httpResponse)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonData)

}
