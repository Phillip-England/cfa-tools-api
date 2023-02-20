package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		res.InvalidRequestMethod(w)
	}

	net.HttpCookie(w, "token", "", -10)
	net.HttpCookie(w, "refresh", "", -10)
	res.Success(w)
}
