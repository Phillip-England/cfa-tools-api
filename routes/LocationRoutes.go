package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/location"
	"github.com/phillip-england/go-http/mid"
)

func LocationRoutes() {

	http.HandleFunc("/location/create",
		mid.CORS(
		mid.Preflight(
		mid.Auth(ctrl.CreateLocation))))

	http.HandleFunc("/location/get",
		mid.CORS(
		mid.Preflight(
		mid.Auth(ctrl.GetLocations))))

		http.HandleFunc("/location/get/",
		mid.CORS(
		mid.Preflight(
		mid.Auth(ctrl.GetLocation))))

}