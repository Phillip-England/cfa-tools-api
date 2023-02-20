package net

import "net/http"

func IsPreflight(w http.ResponseWriter, r *http.Request) (isPreflight bool) {
	return r.Method == "OPTIONS"
}
