package mid

import (
	"errors"
	"net/http"

	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/res"
)

func CSRF(w http.ResponseWriter, r *http.Request) (response func()) {

	type CSRFBody struct {
		CSRF string `json"_csrf"`
	}

	body := CSRFBody{}
	err := lib.GetBody(w, r, &body)
	if err != nil {
		return func() {
			res.ServerError(w, err)
		}
	}

	randomRunes := []rune(body.CSRF)
	err = errors.New("invalid csrf token")
	if string(randomRunes[43]) != "G" {
		return func() {
			res.Forbidden(w)
		}
	}
	if string(randomRunes[55]) != "O" {
		return func() {
			res.Forbidden(w)
		}
	}
	return nil
}
