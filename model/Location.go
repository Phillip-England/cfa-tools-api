package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User      primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name      string             `json:"name" bson:"name"`
	Number    string             `json:"number" bson:"number"`
}

type LocationResponse struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User   primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Number string             `json:"number" bson:"number"`
}

func BuildLocation(userID primitive.ObjectID, name string, number string) (location Location, err error) {
	location = Location{
		User:   userID,
		Name:   name,
		Number: number,
	}
	location.Timestamp()
	return location, nil
}

func (v *Location) Timestamp() {
	now := time.Now()
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	v.UpdatedAt = now
}
