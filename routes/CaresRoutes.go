package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/cares"
	"github.com/phillip-england/go-http/mid"
	"github.com/phillip-england/go-http/model"
)

func CaresRoutes(db model.Db) {

	http.HandleFunc("/cares/delete/", mid.Handler(ctrl.DeleteCares, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "DELETE",
		Auth:      true,
		Location:  true,
	}))

	http.HandleFunc("/cares/create", mid.Handler(ctrl.CreateCares, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "POST",
		Auth:      true,
		Location:  true,
	}))
}
