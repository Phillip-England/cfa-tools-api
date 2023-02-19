package net

import (
	"io"
	"net/http"
)

func GetBody(w http.ResponseWriter, r *http.Request) (body []byte, err error) {
	body, err = io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}