package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/mongo"
)

func LogoutUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {

	lib.HttpCookie(w, "token", "", -10)
	lib.HttpCookie(w, "refresh", "", -10)
	res.Success(w)
}
