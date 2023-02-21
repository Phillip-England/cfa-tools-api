package net

import "net/http"

func ServerStaticFiles() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}