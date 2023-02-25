package ctrl

import (
	"log"
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteCares(w http.ResponseWriter, r *http.Request, db model.Db) {

	id := net.GetURLParam(r.URL.Path)
	caresID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res.ResourceNotFound(w)
		return
	}

	const userKey model.ContextKey = "user"
	user := r.Context().Value(userKey).(model.User)

	log.Println(caresID, user)

}
