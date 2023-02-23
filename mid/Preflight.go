package mid

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/res"
)

func Preflight(w http.ResponseWriter, r *http.Request) (response func()) {
	isPreflight := net.IsPreflight(w, r)
	if isPreflight {
		return func() {
			res.Success(w)
		}
	}
	return nil
}
