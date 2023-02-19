package ctrl

import (
	"net/http"
	"time"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != "PUT" {
		res.InvalidRequestMethod(w)
		return
	}
	
	type requestBody struct {
		CurrentPassword string `json:"current_password"`
		NewPassword string `json:"new_password"`
	}

	body := requestBody{}
	err := net.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}
	
	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	password, err := user.GetDecryptedPassword()
	if err != nil {
		res.ServerError(w, err)
		return
	}

	if body.CurrentPassword != string(password) {
		res.BadReqeust(w, "current password invalid")
		return
	}

	encryptedPassword, err := lib.Encrypt([]byte(body.NewPassword))
	if err != nil {
		res.ServerError(w, err)
		return
	}

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "users")

	filter := bson.D{{
		Key:"$set", Value: bson.D{
			{Key: "password", Value: string(encryptedPassword)},
			{Key: "updated_at", Value: time.Now()},
		},
	}}
	_, err = coll.UpdateByID(ctx, user.ID, filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)

}