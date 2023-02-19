package mid

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, tokenErr := r.Cookie("token")
		refresh, refreshErr := r.Cookie("refresh")

		// found no tokens
		if tokenErr != nil && refreshErr != nil {
			net.Unauthorized(w)
			return
		}

		// found token but not refresh
		if tokenErr == nil && refreshErr != nil {
			jwtData, err := lib.DecodeJWT(token.Value)
			if err != nil {
				net.ServerError(w, err)
				return
			}
			signedString, err := lib.GetJWT(jwtData.(string))
			if err != nil {
				net.ServerError(w, err)
				return
			}
			net.HttpCookie(w, "refresh", signedString, 45)
			ctx, client, disconnect := db.Connect()
			defer disconnect()
			coll := db.Collection(client, "users")
			objectID, err := primitive.ObjectIDFromHex(jwtData.(string))
			if err != nil {
				net.ServerError(w, err)
				return
			}
			var user model.User
			filter := bson.D{{Key: "_id", Value: objectID}}
			err = coll.FindOne(ctx, filter).Decode(&user)
			if err != nil {
				net.ServerError(w, err)
			}
			const userKey model.ContextKey = "user"
			ctx = context.WithValue(r.Context(), userKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		// found refresh but not token
		if tokenErr != nil && refreshErr == nil {
			jwtData, err := lib.DecodeJWT(refresh.Value)
			if err != nil {
				net.ServerError(w, err)
				return
			}
			signedString, err := lib.GetJWT(jwtData.(string))
			if err != nil {
				net.ServerError(w, err)
				return
			}
			net.HttpCookie(w, "token", signedString, 30)
			net.HttpCookie(w, "refresh", refresh.Value, 45)
			ctx, client, disconnect := db.Connect()
			defer disconnect()
			coll := db.Collection(client, "users")
			objectID, err := primitive.ObjectIDFromHex(jwtData.(string))
			if err != nil {
				net.ServerError(w, err)
				return
			}
			var user model.User
			filter := bson.D{{Key: "_id", Value: objectID}}
			err = coll.FindOne(ctx, filter).Decode(&user)
			if err != nil {
				net.ServerError(w, err)
				return
			}
			const userKey model.ContextKey = "user"
			ctx = context.WithValue(r.Context(), userKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		// found both
		if tokenErr == nil && refreshErr == nil {
			jwtData, err := lib.DecodeJWT(token.Value)
			if err != nil {
				net.ServerError(w, err)
				return
			}
			ctx, client, disconnect := db.Connect()
			defer disconnect()
			coll := db.Collection(client, "users")
			objectID, err := primitive.ObjectIDFromHex(jwtData.(string))
			if err != nil {
				net.ServerError(w, err)
				return
			}
			var user model.User
			filter := bson.D{{Key: "_id", Value: objectID}}
			err = coll.FindOne(ctx, filter).Decode(&user)
			if err != nil {
				net.ServerError(w, err)
				return
			}
			const userKey model.ContextKey = "user"
			ctx = context.WithValue(r.Context(), userKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		



		

	}
}