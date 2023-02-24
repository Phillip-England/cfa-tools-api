package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
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

	const locationKey model.ContextKey = "location"
	location := r.Context().Value(locationKey).(model.Location)

	ctx, client, disconnect := db.Connect()
	defer disconnect()

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
