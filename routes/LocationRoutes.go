package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/location"
	"github.com/phillip-england/go-http/mid"
	"github.com/phillip-england/go-http/model"
)

func LocationRoutes(db model.Db) {

	http.HandleFunc("/location/create", mid.Handler(ctrl.CreateLocation, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "POST",
		Auth:      true,
	}))

	http.HandleFunc("/location/get", mid.Handler(ctrl.GetLocations, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/location/get/", mid.Handler(ctrl.GetLocation, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/location/delete/", mid.Handler(ctrl.DeleteLocation, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "DELETE",
		Auth:      true,
	}))

	http.HandleFunc("/location/update/", mid.Handler(ctrl.UpdateLocation, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "PUT",
		Auth:      true,
	}))

	http.HandleFunc("/location/select/", mid.Handler(ctrl.SelectLocation, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/location/get/active", mid.Handler(ctrl.GetActiveLocation, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/location/drop/active", mid.Handler(ctrl.DropActiveLocation, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

}
