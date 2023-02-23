package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/location"
	"github.com/phillip-england/go-http/mid"
)

func LocationRoutes() {
	http.HandleFunc("/location/create", mid.Handler("POST", ctrl.CreateLocation, mid.MidOptionsUser()))
	http.HandleFunc("/location/get", mid.Handler("GET", ctrl.GetLocations, mid.MidOptionsUser()))
	http.HandleFunc("/location/get/", mid.Handler("GET", ctrl.GetLocation, mid.MidOptionsUser()))
	http.HandleFunc("/location/delete/", mid.Handler("DELETE", ctrl.DeleteLocation, mid.MidOptionsUser()))
	http.HandleFunc("/location/update/", mid.Handler("PUT", ctrl.UpdateLocation, mid.MidOptionsUser()))
	http.HandleFunc("/location/select/", mid.Handler("GET", ctrl.SelectLocation, mid.MidOptionsUser()))
	http.HandleFunc("/location/get/active", mid.Handler("GET", ctrl.GetActiveLocation, mid.MidOptionsUser()))
	http.HandleFunc("/location/drop/active", mid.Handler("GET", ctrl.DropActiveLocation, mid.MidOptionsUser()))
}
