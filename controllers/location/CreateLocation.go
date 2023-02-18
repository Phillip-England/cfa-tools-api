package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
)

func CreateLocation(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		net.MessageResponse(w, "invalid request method", 400)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	body := net.GetBody(w, r)
	var location model.Location
	err := json.Unmarshal(body, &location)
	if err != nil {
		net.ServerError(w)
		return
	}
	location.Timestamp()
	location.User = user.ID

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "locations")

	_, err = coll.InsertOne(ctx, location)
	if err != nil {
		net.ServerError(w)
		return
	}

	net.MessageResponse(w, "success", 200)
}