package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/cares"
	"github.com/phillip-england/go-http/mid"
)

func CaresRoutes() {

	http.HandleFunc("/cares/delete/", mid.Handler(ctrl.DeleteCares, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "DELETE",
		Auth: true,
		Location: true,
	}))

	http.HandleFunc("/cares/create", mid.Handler(ctrl.CreateCares, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "POST",
		Auth: true,
		Location: true,
	}))
}
