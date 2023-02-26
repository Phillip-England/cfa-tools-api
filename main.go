package main

import (
	"github.com/joho/godotenv"
	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/lib"
	"github.com/phillip-england/go-http/routes"
)

func main() {
	godotenv.Load()
	client, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Disconnect(client)
	routes.Mount(client)
	lib.Serve()
}
