package ctrl

import (
	"net/http"
	"strings"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)




func LoginUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		net.MessageResponse(w, "invalid request method", 400)
		return
	}

	type requestBody struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	body := requestBody{}
	err := net.GetBody(w, r, &body)
	if err != nil {
		net.ServerError(w, err)
		return
	}

	body.Email = strings.ToLower(body.Email)

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "users")

	var userExists model.User
	filter := bson.D{{Key: "email", Value: body.Email}}
	err = coll.FindOne(ctx, filter).Decode(&userExists)

	if err == mongo.ErrNoDocuments {
		net.MessageResponse(w, "invalid credentials", 400)
		return
	} else if err != nil {
		net.ServerError(w, err)
		return
	}

	decryptedPasswordBytes, err := lib.Decrypt([]byte(userExists.Password))
	if err != nil {
		net.ServerError(w, err)
		return
	}

	if body.Password != string(decryptedPasswordBytes) {
		net.MessageResponse(w, "invalid credentials", 400)
		return
	}

	signedString, err := lib.GetJWT(userExists.ID.Hex())
	if err != nil {
		net.ServerError(w, err)
		return
	}

	net.HttpCookie(w, "token", signedString, 30)

	net.MessageResponse(w, "success", 200)

}