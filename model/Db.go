package model

import (
	"context"

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
