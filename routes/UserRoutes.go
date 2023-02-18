package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/controllers/user"
	mid "github.com/phillip-england/go-http/middleware"
)

func UserRoutes() {

	http.HandleFunc("/user/create", 
		mid.CORS(
		mid.Preflight(ctrl.CreateUser)))

	http.HandleFunc("/user/login", 
		mid.CORS(
		mid.Preflight(ctrl.LoginUser)))

	http.HandleFunc("/user/deleteall", 
		mid.CORS(ctrl.DeleteAllUsers))

	http.HandleFunc("/user/get",
		mid.CORS(
		mid.Preflight(
		mid.Auth(ctrl.GetUser))))	

	http.HandleFunc("/user/logout",
		mid.CORS(
		mid.Preflight(
		mid.Auth(ctrl.LogoutUser))))
	

}