package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
)

func CreateLocation(w http.ResponseWriter, r *http.Request, db model.Db) {

	type requestBody struct {
		Name   string `json:"name"`
		Number string `json:"number"`
		CSRF   string `json:"_csrf"`
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

	location, err := model.BuildLocation(user.ID, body.Name, body.Number)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	coll := db.Collection("locations")

	_, err = coll.InsertOne(db.Ctx, location)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)
}
