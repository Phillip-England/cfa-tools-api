package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/cares"
	"github.com/phillip-england/go-http/mid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CaresRoutes(client *mongo.Client) {

	http.HandleFunc("/cares/delete/", mid.Handler(ctrl.DeleteCares, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "DELETE",
		Auth:      true,
		Location:  true,
	}))

	http.HandleFunc("/cares/create", mid.Handler(ctrl.CreateCares, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "POST",
		CSRF:      true,
		Auth:      true,
		Location:  true,
	}))

	http.HandleFunc("/cares/get/all", mid.Handler(ctrl.GetAllCares, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
		Location:  true,
	}))

	http.HandleFunc("/cares/update/", mid.Handler(ctrl.UpdateCares, client, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "PUT",
		CSRF:      true,
		Auth:      true,
		Location:  true,
	}))

}
