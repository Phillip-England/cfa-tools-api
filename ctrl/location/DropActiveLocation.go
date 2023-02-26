package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/mongo"
)

func DropActiveLocation(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	lib.HttpCookie(w, "location", "", -10)
	res.Success(w)

}
