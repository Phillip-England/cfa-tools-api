package mid

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/model"
	"github.com/phillip-england/go-http/res"
)

func Handler(controller func(w http.ResponseWriter, r *http.Request, db model.Db), db model.Db, options Options) http.HandlerFunc {
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

		if r.Method != options.Method {
			res.InvalidRequestMethod(w)
			return
		}

		if options.Auth {
			var response func()
			ctx, response = Auth(w, r, db)
			if response != nil {
				response()
				return
			}
		}

		if options.Location {
			var response func()
			ctx, response = Location(ctx, w, r, db)
			if response != nil {
				response()
				return
			}

		}

		if ctx == nil {
			controller(w, r, db)
			return
		}

		controller(w, r.WithContext(ctx), db)

	}

}
