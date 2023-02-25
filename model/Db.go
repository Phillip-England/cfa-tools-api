package model

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type Db struct {
	Ctx    context.Context
	Cancel context.CancelFunc
	Client *mongo.Client
}

func (v *Db) CloseConnection() {
	v.Cancel()
	if err := v.Client.Disconnect(v.Ctx); err != nil {
		panic(err)
	}
}

func (v *Db) Collection(name string) (coll *mongo.Collection) {
	coll = v.Client.Database(os.Getenv("MONGO_DB")).Collection(name)
	return coll
}
