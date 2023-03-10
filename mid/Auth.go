package mid

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

func Auth(client *mongo.Client, w http.ResponseWriter, r *http.Request) (httpctx context.Context, response func()) {


	token, tokenErr := r.Cookie("token")
	refresh, refreshErr := r.Cookie("refresh")
	var jwtData interface{}
	var err error

	if tokenErr != nil && refreshErr != nil {
		return nil, func() {
			res.Unauthorized(w)
		}
	}
	
	if tokenErr == nil && refreshErr != nil || tokenErr == nil && refreshErr == nil {
		jwtData, err = lib.DecodeJWT(token.Value)
		if err != nil {
			return nil, func() {
				res.ServerError(w, err)
			}
		}
	}
	
	
	if tokenErr != nil && refreshErr == nil {
		jwtData, err = lib.DecodeJWT(refresh.Value)
		if err != nil {
			return nil, func() {
				res.ServerError(w, err)
			}
		}
	}

	signedStringRefresh, err := lib.GetJWT(jwtData.(string))
	if err != nil {
		return nil, func() {
			res.ServerError(w, err)
		}
	}
	signedStringToken, err := lib.GetJWT(jwtData.(string))
	if err != nil {
		return nil, func() {
			res.ServerError(w, err)
		}
	}
	lib.HttpCookie(w, "refresh", signedStringRefresh, 45)
	lib.HttpCookie(w, "token", signedStringToken, 30)
	coll := db.Collection(client, "users")
	objectID, err := primitive.ObjectIDFromHex(jwtData.(string))
	if err != nil {
		return nil, func() {
			res.ServerError(w, err)
		}
	}
	var user model.User
	filter := bson.D{{Key: "_id", Value: objectID}}
	err = coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, func() {
			res.ServerError(w, err)
		}
	}
	const userKey model.ContextKey = "user"
	httpctx = context.WithValue(r.Context(), userKey, user)

	return httpctx, nil

}
