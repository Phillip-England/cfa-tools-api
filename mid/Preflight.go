package mid

import (
	"net/http"

	"github.com/phillip-england/go-http/res"
)

func Preflight(w http.ResponseWriter, r *http.Request) (response func()) {
	if r.Method == "OPTIONS" {
		return func() {
			res.Success(w)
		}
	}
	return nil
}
