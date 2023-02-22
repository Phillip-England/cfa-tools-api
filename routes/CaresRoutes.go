package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/cares"
	"github.com/phillip-england/go-http/mid"
)

func CaresRoutes() {

	http.HandleFunc("/cares/create",
		mid.CORS(
			mid.Preflight(
				mid.Auth(ctrl.CreateCares))))

	http.HandleFunc("/cares/delete/",
		mid.CORS(
			mid.Preflight(
				mid.Auth(ctrl.DeleteCares))))

}
