package ctrl

import (
	"log"
	"net/http"

	"github.com/phillip-england/go-http/net"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != "PUT" {
		net.InvalidRequestMethod(w)
		return
	}
	
	type requestBody struct {
		Email string `json:"email"`
		CurrentPassword string `json:"current_password"`
		NewPassword string `json:"new_password"`
	}

	// const userKey model.ContextKey = "user"
	// user := r.Context().Value(userKey).(model.User)

	body := requestBody{}
	err := net.GetBody(w, r, &body)
	if err != nil {
		net.ServerError(w, err)
		return
	}

	log.Println(body)

	// var updatedUser model.UpdatedUser
	// err = json.Unmarshal(body, &updatedUser)
	// if err != nil {
	// 	net.ServerError(w, err)
	// 	return
	// }
	// updatedUser.Timestamp()

	// password, err := lib.Decrypt([]byte(user.Password))
	// if err != nil {
	// 	net.ServerError(w, err)
	// 	return
	// }

	// if updatedUser.CurrentPassword != string(password) {
	// 	net.BadReqeust(w, "current password invalid")
	// 	return
	// }

	// encryptedPassword, err := lib.Encrypt([]byte(updatedUser.NewPassword))
	// if err != nil {
	// 	net.ServerError(w, err)
	// 	return
	// }

	// ctx, client, disconnect := db.Connect()
	// defer disconnect()
	// coll := db.Collection(client, "users")

	// filter := bson.D{{Key: "password", Value: encryptedPassword}}
	// _, err = coll.UpdateByID(ctx, user.ID, filter)
	// if err != nil {
	// 	net.ServerError(w, err)
	// 	return
	// }

	net.Success(w)

}