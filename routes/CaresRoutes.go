package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/cares"
	"github.com/phillip-england/go-http/mid"
)

func CaresRoutes() {
	http.HandleFunc("/cares/delete/", mid.Handler("DELETE", ctrl.DeleteCares, mid.MidOptionsUser()))
	http.HandleFunc("/cares/create", mid.Handler("POST", ctrl.CreateCares, mid.MidOptionsUser()))
}
