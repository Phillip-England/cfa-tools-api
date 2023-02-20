package net

import (
	"net/http"
	"time"
)

func HttpCookie(w http.ResponseWriter, name string, value string, expireMinutes int) {
	minutes := time.Duration(expireMinutes)
	expire := time.Now().Add(minutes * time.Minute)
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expire,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}
