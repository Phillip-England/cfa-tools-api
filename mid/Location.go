package mid

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Location(ctx context.Context, w http.ResponseWriter, r *http.Request) (newctx context.Context, response func()) {

	const userKey model.ContextKey = "user"
	const locationKey model.ContextKey = "location"
	user := ctx.Value(userKey).(model.User)

	cookie, err := r.Cookie("location")
	if err != nil {
		return nil, func() {
			res.BadReqeust(w, "no active location")
		}
	}

	locationID, err := primitive.ObjectIDFromHex(cookie.Value)
	if err != nil {
		res.ServerError(w, err)
		return
	}

	ctx, client, disconnect := db.Connect()
	defer disconnect()
	coll := db.Collection(client, "locations")

	filter := bson.D{{Key: "_id", Value: locationID}}
	var location model.Location
	err = coll.FindOne(ctx, filter).Decode(&location)
	if err == mongo.ErrNoDocuments {
		return nil, func() {
			res.ResourceNotFound(w)
		}
	} else if err != nil {
		return nil, func() {
			res.ServerError(w, err)
		}
	}

	if location.User != user.ID {
		return nil, func() {
			res.Forbidden(w)
		}
	}

	newctx = context.WithValue(r.Context(), userKey, user)
	newctx = context.WithValue(newctx, locationKey, location)

	return newctx, nil

}