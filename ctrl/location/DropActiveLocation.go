package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
)

func DropActiveLocation(w http.ResponseWriter, r *http.Request) {

	net.HttpCookie(w, "location", "", -10)
	res.Success(w)

}
