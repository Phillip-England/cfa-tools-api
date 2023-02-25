package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request, db model.Db) {

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

	user, err := model.BuildUser(body.Email, body.Password)
	if err != nil {
		res.ServerError(w, err)
	}

	coll := db.Collection("users")

	var userExists model.User
	filter := bson.D{{Key: "email", Value: user.Email}}
	err = coll.FindOne(db.Ctx, filter).Decode(&userExists)
	if userExists.Email == user.Email && err != mongo.ErrNoDocuments {
		res.MessageResponse(w, "user already exists", 400)
		return
	}
	if err != mongo.ErrNoDocuments && err != nil {
		res.ServerError(w, err)
		return
	}

	result, err := coll.InsertOne(db.Ctx, user)
	if err != nil {
		res.ServerError(w, err)
		return
	}
	id := result.InsertedID.(primitive.ObjectID)

	signedString, err := lib.GetJWT(id.Hex())
	if err != nil {
		res.ServerError(w, err)
		return
	}

	net.HttpCookie(w, "token", signedString, 30)

	res.Success(w)

}
