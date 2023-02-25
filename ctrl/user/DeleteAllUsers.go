package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteAllUsers(w http.ResponseWriter, r *http.Request, db model.Db) {

	coll := db.Collection("users")
	coll.DeleteMany(db.Ctx, bson.D{})

	coll = db.Collection("locations")
	coll.DeleteMany(db.Ctx, bson.D{})

	coll = db.Collection("cares")
	coll.DeleteMany(db.Ctx, bson.D{})

	res.Success(w)

}
