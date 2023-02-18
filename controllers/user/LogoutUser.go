package ctrl

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	net.HttpCookie(w, "token", "", -10)
	net.HttpCookie(w, "refresh", "", -10)
	net.Success(w)
}