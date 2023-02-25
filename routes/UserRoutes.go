package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/user"
	"github.com/phillip-england/go-http/mid"
	"github.com/phillip-england/go-http/model"
)

func UserRoutes(db model.Db) {

	http.HandleFunc("/user/create", mid.Handler(ctrl.CreateUser, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "POST",
	}))

	http.HandleFunc("/user/login", mid.Handler(ctrl.LoginUser, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "POST",
	}))

	http.HandleFunc("/user/deleteall", mid.Handler(ctrl.DeleteAllUsers, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "DELETE",
	}))

	http.HandleFunc("/user/get", mid.Handler(ctrl.GetUser, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/user/logout", mid.Handler(ctrl.LogoutUser, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "GET",
		Auth:      true,
	}))

	http.HandleFunc("/user/update/password", mid.Handler(ctrl.UpdateUserPassword, db, mid.Options{
		CORS:      true,
		Preflight: true,
		Method:    "PUT",
		Auth:      true,
	}))

}
