package mid

import (
	"context"
	"net/http"

	"github.com/phillip-england/go-http/res"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handler(controller func(client *mongo.Client, w http.ResponseWriter, r *http.Request), client *mongo.Client, options Options) http.HandlerFunc {
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
			ctx, response = Auth(client, w, r)
			if response != nil {
				response()
				return
			}
		}

		if options.Location {
			var response func()
			ctx, response = Location(ctx, client, w, r)
			if response != nil {
				response()
				return
			}

		}

		if ctx == nil {
			controller(client, w, r)
			return
		}

		controller(client, w, r.WithContext(ctx))

	}

}
