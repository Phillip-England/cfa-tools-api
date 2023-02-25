package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/user"
	"github.com/phillip-england/go-http/mid"
)

func UserRoutes() {

	http.HandleFunc("/user/create", mid.Handler(ctrl.CreateUser, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "POST",
	}))

	http.HandleFunc("/user/login", mid.Handler(ctrl.LoginUser, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "POST",
	}))

	http.HandleFunc("/user/deleteall", mid.Handler(ctrl.DeleteAllUsers, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "DELETE",
	}))

	http.HandleFunc("/user/get", mid.Handler(ctrl.GetUser, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

	http.HandleFunc("/user/logout", mid.Handler(ctrl.LogoutUser, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "GET",
		Auth: true,
	}))

	http.HandleFunc("/user/update/password", mid.Handler(ctrl.UpdateUserPassword, mid.Options{
		CORS: true,
		Preflight: true,
		Method: "PUT",
		Auth: true,
	}))

	

}
