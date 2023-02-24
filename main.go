package main

import (
	"github.com/joho/godotenv"
	"github.com/phillip-england/go-http/db"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/routes"
)

func main() {
	db := db.Connect()
	defer db.CloseConnection()
	godotenv.Load()
	routes.Mount(db)
	net.Serve()
}
