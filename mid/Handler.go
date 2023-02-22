package mid

import (
	"net/http"

	"github.com/phillip-england/go-http/net"
)

func Handler(endpoint string, controller func(w http.ResponseWriter, r *http.Request), options Options) (next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {

		if options.CORS {
			net.SetCORS(w)
		}
		
		
	}


}