package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/location"
	"github.com/phillip-england/go-http/mid"
	"go.mongodb.org/mongo-driver/mongo"
)

func LocationRoutes(client *mongo.Client) {

	http.HandleFunc("/location/create", mid.Handler(ctrl.CreateLocation, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "POST",
		Auth:      true,
	}))

	http.HandleFunc("/location/get", mid.Handler(ctrl.GetLocations, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/location/get/", mid.Handler(ctrl.GetLocation, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/location/delete/", mid.Handler(ctrl.DeleteLocation, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "DELETE",
		Auth:      true,
	}))

	http.HandleFunc("/location/update/", mid.Handler(ctrl.UpdateLocation, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "PUT",
		Auth:      true,
	}))

	http.HandleFunc("/location/select/", mid.Handler(ctrl.SelectLocation, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/location/get/active", mid.Handler(ctrl.GetActiveLocation, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/location/drop/active", mid.Handler(ctrl.DropActiveLocation, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

}
