package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		net.MessageResponse(w, "invalid request method", 400)
	}

	net.HttpCookie(w, "token", "", -10)
	net.HttpCookie(w, "refresh", "", -10)
	net.Success(w)
}