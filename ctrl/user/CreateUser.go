package ctrl

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		net.MessageResponse(w, "invalid request method", http.StatusBadRequest)
		return
	}

	body, err := net.GetBody(w, r)
	if err != nil {
		net.ServerError(w, err)
		return
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		net.ServerError(w, err)
		return
	}

	user.Timestamp()
	user.Email = strings.ToLower(user.Email)
	encryptedPassword, err := lib.Encrypt([]byte(user.Password))
	if err != nil {
		net.ServerError(w, err)
		return
	}
	user.Password = string(encryptedPassword)
	
	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "users")
	
	var userExists model.User
	filter := bson.D{{Key: "email", Value: user.Email}}
	err = coll.FindOne(ctx, filter).Decode(&userExists)
	if userExists.Email == user.Email && err != mongo.ErrNoDocuments {
		net.MessageResponse(w, "user already exists", 400)
		return
	}
	if err != mongo.ErrNoDocuments && err != nil {
		net.ServerError(w, err)
		return
	}

	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		net.ServerError(w, err)
		return
	}
	id := res.InsertedID.(primitive.ObjectID)

	signedString, err := lib.GetJWT(id.Hex())
	if err != nil {
		net.ServerError(w, err)
		return
	}

	net.HttpCookie(w, "token", signedString, 30)
	
	net.MessageResponse(w, "success", 200)

}