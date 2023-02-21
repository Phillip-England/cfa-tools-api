package mid

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, tokenErr := r.Cookie("token")
		refresh, refreshErr := r.Cookie("refresh")
		var jwtData interface{}
		var err error

		if tokenErr != nil && refreshErr != nil {
			res.Unauthorized(w)
			return
		}

		if tokenErr == nil && refreshErr != nil || tokenErr == nil && refreshErr == nil {
			jwtData, err = lib.DecodeJWT(token.Value)
			if err != nil {
				res.ServerError(w, err)
				return
			}
		}

		if tokenErr != nil && refreshErr == nil {
			jwtData, err = lib.DecodeJWT(refresh.Value)
			if err != nil {
				res.ServerError(w, err)
				return
			}
		}

		signedStringRefresh, err := lib.GetJWT(jwtData.(string))
		if err != nil {
			res.ServerError(w, err)
			return
		}
		signedStringToken, err := lib.GetJWT(jwtData.(string))
		if err != nil {
			res.ServerError(w, err)
			return
		}
		net.HttpCookie(w, "refresh", signedStringRefresh, 45)
		net.HttpCookie(w, "token", signedStringToken, 30)
		ctx, client, disconnect := db.Connect()
		defer disconnect()
		coll := db.Collection(client, "users")
		objectID, err := primitive.ObjectIDFromHex(jwtData.(string))
		if err != nil {
			res.ServerError(w, err)
			return
		}
		var user model.User
		filter := bson.D{{Key: "_id", Value: objectID}}
		err = coll.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			res.ServerError(w, err)
		}
		const userKey model.ContextKey = "user"
		ctx = context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
