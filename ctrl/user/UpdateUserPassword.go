package ctrl

import (
	"log"
	"net/http"
	"time"

	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUserPassword(w http.ResponseWriter, r *http.Request, db model.Db) {

	type requestBody struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
		CSRF            string `json:"_csrf"`
	}

	body := requestBody{}
	err := net.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	err = net.IsCSRF(body.CSRF)
	if err != nil {
		log.Println(err)
		res.Forbidden(w)
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

	coll := db.Collection("users")

	filter := bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "password", Value: string(encryptedPassword)},
			{Key: "updated_at", Value: time.Now()},
		},
	}}
	_, err = coll.UpdateByID(db.Ctx, user.ID, filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)

}
