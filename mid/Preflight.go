package mid

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
)

func Preflight(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isPreflight := net.IsPreflight(w, r)
		if isPreflight {
			res.Success(w)
		} else {
			next(w, r)
		}
	}
}