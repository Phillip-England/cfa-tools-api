package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
)

func DropActiveLocation(w http.ResponseWriter, r *http.Request, db model.Db) {

	net.HttpCookie(w, "location", "", -10)
	res.Success(w)

}
