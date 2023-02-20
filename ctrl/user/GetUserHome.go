package ctrl

import (
	"net/http"
	"text/template"

	"github.com/phillip-england/go-http/res"
)

func appHome(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./views/appHome.html")
	if err != nil {
		res.ServerError(w, err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		res.ServerError(w, err)
		return
	}

}