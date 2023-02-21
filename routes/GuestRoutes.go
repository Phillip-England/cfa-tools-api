package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/guest"
	"github.com/phillip-england/go-http/mid"
)

func GuestRoutes() {

	http.HandleFunc("/guest/home",
		mid.CORS(
			mid.Preflight(ctrl.GuestHome)))

}