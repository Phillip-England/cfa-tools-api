package mid

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
)

func Preflight(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isPreflight := net.IsPreflight(w, r)
		if isPreflight {
			net.MessageResponse(w, "success", http.StatusOK)
		} else {
			next(w, r)
		}
	}
}