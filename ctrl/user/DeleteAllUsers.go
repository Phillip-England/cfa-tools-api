package ctrl

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteAllUsers(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	coll := db.Collection(client, "users")
	coll.DeleteMany(context.Background(), bson.D{})

	coll = db.Collection(client, "locations")
	coll.DeleteMany(context.Background(), bson.D{})

	coll = db.Collection(client, "cares")
	coll.DeleteMany(context.Background(), bson.D{})

	res.Success(w)

}
