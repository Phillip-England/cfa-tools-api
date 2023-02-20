package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (ctx context.Context, client *mongo.Client, disconnect func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	disconnect = func() {
		cancel()
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	return ctx, client, disconnect
}
