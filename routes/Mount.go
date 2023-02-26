package routes

import "go.mongodb.org/mongo-driver/mongo"

func Mount(client *mongo.Client) {
	UserRoutes(client)
	LocationRoutes(client)
	CaresRoutes(client)
}
