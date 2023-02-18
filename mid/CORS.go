package mid

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		net.SetCORS(w)
		next(w, r)
	}
}