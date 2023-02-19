package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
)

func CreateLocation(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		net.InvalidRequestMethod(w)
		return
	}
	
	type requestBody struct {
		Name string `json:"name"`
		Number string `json:"number"`
	}
	
	body := requestBody{}
	err := net.GetBody(w, r, &body)
	if err != nil {
		net.ServerError(w, err)
		return
	}
	
	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	location, err := model.BuildLocation(user.ID, body.Name, body.Number)
	if err != nil {
		net.ServerError(w, err)
		return
	}

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "locations")

	_, err = coll.InsertOne(ctx, location)
	if err != nil {
		net.ServerError(w, err)
		return
	}

	net.Success(w)
}