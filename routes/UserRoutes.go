package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/user"
	"github.com/phillip-england/go-http/mid"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(client *mongo.Client) {

	http.HandleFunc("/user/create", mid.Handler(ctrl.CreateUser, client, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "POST",
	}))

	http.HandleFunc("/user/login", mid.Handler(ctrl.LoginUser, client, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "POST",
	}))

	http.HandleFunc("/user/deleteall", mid.Handler(ctrl.DeleteAllUsers, client, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "DELETE",
	}))

	http.HandleFunc("/user/get", mid.Handler(ctrl.GetUser, client, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

	http.HandleFunc("/user/logout", mid.Handler(ctrl.LogoutUser, client, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

	http.HandleFunc("/user/update/password", mid.Handler(ctrl.UpdateUserPassword, client, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "PUT",
		Auth: true,
	}))

}
