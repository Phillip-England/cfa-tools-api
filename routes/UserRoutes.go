package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/user"
	"github.com/phillip-england/go-http/mid"
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