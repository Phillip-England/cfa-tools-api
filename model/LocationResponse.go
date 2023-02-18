package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationResponse struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User 			primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Name     string `json:"name" bson:"name"`
	Number  string `json:"number" bson:"number"`
}
