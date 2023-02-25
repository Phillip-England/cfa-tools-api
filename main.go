package main

import (
	"github.com/joho/godotenv"
	"github.com/phillip-england/go-http/net"
	"github.com/phillip-england/go-http/routes"
)

func main() {
	godotenv.Load()
	db := net.ConnectDb()
	defer db.CloseConnection()
	routes.Mount(db)
	net.Serve()
}
