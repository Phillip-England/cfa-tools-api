package net

import (
	"context"
	"os"
	"time"

	"github.com/phillip-england/go-http/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() (db model.Db) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		cancel()
	}
	db = model.Db{
		Ctx:    ctx,
		Cancel: cancel,
		Client: client,
	}
	return db
}
