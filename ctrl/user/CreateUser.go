package ctrl

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	body := requestBody{}
	err := lib.GetBody(w, r, &body)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	if body.Password != body.ConfirmPassword {
		res.BadReqeust(w, "passwords must match")
		return
	}

	user, err := model.BuildUser(body.Email, body.Password)
	if err != nil {
		res.ServerError(w, err)
	}

	err = user.Validate()
	if err != nil {
		res.MessageResponse(w, err.Error(), 400)
		return
	}

	err = user.EncryptPassword()
	if err != nil {
		res.ServerError(w, err)
		return
	}

	coll := db.Collection(client, "users")

	var userExists model.User
	filter := bson.D{{Key: "email", Value: user.Email}}
	err = coll.FindOne(context.Background(), filter).Decode(&userExists)
	if userExists.Email == user.Email && err != mongo.ErrNoDocuments {
		res.MessageResponse(w, "User already exists", 400)
		return
	}
	if err != mongo.ErrNoDocuments && err != nil {
		res.ServerError(w, err)
		return
	}

	result, err := coll.InsertOne(context.Background(), user)
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

	lib.HttpCookie(w, "token", signedString, 30)

	res.Success(w)

}
