package ctrl

import (
	"net/http"
	"strings"

	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginUser(w http.ResponseWriter, r *http.Request, db model.Db) {

	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	body := requestBody{}
	err := net.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	body.Email = strings.ToLower(body.Email)

	coll := db.Collection("users")

	var userExists model.User
	filter := bson.D{{Key: "email", Value: body.Email}}
	err = coll.FindOne(db.Ctx, filter).Decode(&userExists)

	if err == mongo.ErrNoDocuments {
		res.BadReqeust(w, "invalid credentials")
		return
	} else if err != nil {
		res.ServerError(w, err)
		return
	}

	decryptedPasswordBytes, err := lib.Decrypt([]byte(userExists.Password))
	if err != nil {
		res.ServerError(w, err)
		return
	}

	if body.Password != string(decryptedPasswordBytes) {
		res.BadReqeust(w, "invalid credentials")
		return
	}

	signedString, err := lib.GetJWT(userExists.ID.Hex())
	if err != nil {
		res.ServerError(w, err)
		return
	}

	net.HttpCookie(w, "token", signedString, 30)

	res.MessageResponse(w, "success", 200)

}
