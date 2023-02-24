package mid

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/res"
)

func Handler(method string, controller func(w http.ResponseWriter, r *http.Request),  options Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var ctx context.Context = nil

		if options.CORS {
			CORS(w)
		}

		if options.Preflight {
			response := Preflight(w, r)
			if response != nil {
				response()
				return
			}
		}

		if r.Method != method {
			res.InvalidRequestMethod(w)
			return
		}

		if options.Auth {
			var response func()
			ctx, response = Auth(w, r)
			if response != nil {
				response()
				return
			}
		}

		if options.Location {
			var response func()
			ctx, response = Location(ctx, w, r)
			if response != nil {
				response()
				return
			}

		}

		if ctx == nil {
			controller(w, r)
			return
		}

		controller(w, r.WithContext(ctx))

	}

}
