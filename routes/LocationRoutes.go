package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/location"
	"github.com/phillip-england/go-http/mid"
)

func LocationRoutes() {

	http.HandleFunc("/location/create", mid.Handler(ctrl.CreateLocation, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "POST",
		Auth: true,
	}))

	http.HandleFunc("/location/get", mid.Handler(ctrl.GetLocations, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

	http.HandleFunc("/location/get/", mid.Handler(ctrl.GetLocation, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

	http.HandleFunc("/location/delete/", mid.Handler(ctrl.DeleteLocation, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "DELETE",
		Auth: true,
	}))

	http.HandleFunc("/location/update/", mid.Handler(ctrl.UpdateLocation, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "PUT",
		Auth: true,
	}))

	http.HandleFunc("/location/select/", mid.Handler(ctrl.SelectLocation, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

	http.HandleFunc("/location/get/active", mid.Handler(ctrl.GetActiveLocation, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

	http.HandleFunc("/location/drop/active", mid.Handler(ctrl.DropActiveLocation, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

}
