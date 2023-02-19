package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/net"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		net.InvalidRequestMethod(w)		
		return
	}

	ctx, client, disconnect := db.Connect()
	defer disconnect()

	coll := db.Collection(client, "users")
	coll.DeleteMany(ctx, bson.D{})

	coll = db.Collection(client, "locations")
	coll.DeleteMany(ctx, bson.D{})

	net.Success(w)

}