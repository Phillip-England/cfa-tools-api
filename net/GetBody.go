package net

import (
	"io"
	"net/http"
)

func GetBody(w http.ResponseWriter, r *http.Request) (body []byte) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ServerError(w)
	}
	return body
}