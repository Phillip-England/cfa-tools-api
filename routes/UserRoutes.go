package routes

import (
	"net/http"

	ctrl "github.com/phillip-england/go-http/ctrl/user"
	"github.com/phillip-england/go-http/mid"
)

func UserRoutes() {

	http.HandleFunc("/user/create", mid.Handler("POST", ctrl.CreateUser, mid.MidOptionsGuest()))
	http.HandleFunc("/user/login", mid.Handler("POST", ctrl.LoginUser, mid.MidOptionsGuest()))
	http.HandleFunc("/user/deleteall", mid.Handler("DELETE", ctrl.DeleteAllUsers, mid.MidOptionsGuest()))
	http.HandleFunc("/user/get", mid.Handler("GET", ctrl.GetUser, mid.MidOptionsUser()))
	http.HandleFunc("/user/logout", mid.Handler("GET", ctrl.LogoutUser, mid.MidOptionsUser()))
	http.HandleFunc("/user/update/password", mid.Handler("PUT", ctrl.UpdateUserPassword, mid.MidOptionsUser()))

}
