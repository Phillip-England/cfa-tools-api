package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/controllers/location"
	mid "github.com/phillip-england/go-http/middleware"
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