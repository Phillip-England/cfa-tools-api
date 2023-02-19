package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		net.InvalidRequestMethod(w)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	body, err := net.GetBody(w, r)
	if err != nil {
		net.ServerError(w, err)
		return
	}

	var updatedUser model.UpdatedUser
	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		net.ServerError(w, err)
		return
	}
	updatedUser.Timestamp()

	password, err := lib.Decrypt([]byte(user.Password))
	if err != nil {
		net.ServerError(w, err)
		return
	}

	if updatedUser.CurrentPassword != string(password) {
		net.BadReqeust(w, "current password invalid")
		return
	}

	encryptedPassword, err := lib.Encrypt([]byte(updatedUser.NewPassword))
	if err != nil {
		net.ServerError(w, err)
		return
	}

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "users")

	filter := bson.D{{Key: "password", Value: encryptedPassword}}
	_, err = coll.UpdateByID(ctx, user.ID, filter)
	if err != nil {
		net.ServerError(w, err)
		return
	}

	net.Success(w)

}