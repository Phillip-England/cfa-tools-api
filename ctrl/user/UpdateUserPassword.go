package ctrl

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUserPassword(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
		CSRF            string `json:"_csrf"`
	}

	body := requestBody{}
	err := lib.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	err = lib.IsCSRF(body.CSRF)
	if err != nil {
		log.Println(err)
		res.Forbidden(w)
		return
	}

	user := r.Context().Value(model.GetUserKey()).(model.User)

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

	coll := db.Collection(client, "users")

	filter := bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "password", Value: string(encryptedPassword)},
			{Key: "updated_at", Value: time.Now()},
		},
	}}
	_, err = coll.UpdateByID(context.Background(), user.ID, filter)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	res.Success(w)

}
