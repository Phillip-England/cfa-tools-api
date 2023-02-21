package net

import (
	"log"
	"net/http"
)

func RunServer() {
	log.Println("Serving application on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
