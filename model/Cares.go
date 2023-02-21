package model

import (
	"context"
	"log"
	"time"

	"github.com/phillip-england/go-http/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cares struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User              primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Location          primitive.ObjectID `json:"location,omitempty" bson:"location,omitempty"`
	CreatedAt         time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	GuestName         string             `json:"guestName" bson:"guestName"`
	OrderNumber       string             `json:"orderNumber" bson:"orderNumber"`
	Incident          string             `json:"incident" bson:"incident"`
	ReplacementAction string             `json:"replacementAction" bson:"replacementAction"`
	ReplacementCode   int                `json:"replacementCode,omitempty" bson:"replacementCode,omitempty"`
}

func (v *Cares) Timestamp() {
	now := time.Now()
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	v.UpdatedAt = now
}

func (v *Cares) SetReplacementCode(ctx context.Context, coll *mongo.Collection) (err error) {

	pipeline := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "nil"},
			{Key: "max", Value: bson.D{
				{Key: "$max", Value: "replacementCode"},
			}},
		}},
	}

	// pipeline := bson.D{{Key: "$max", Value: "replacementCode"}}

	var result []Cares
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{pipeline})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if cursor.All(ctx, &result); err != nil {
		return err
	}
	
	log.Println(result)

	if len(result) == 0 {
		v.ReplacementCode = 1000
		return nil
	}

	lib.PrintType(result)
	// v.ReplacementCode = result[0]["maxReplacementCode"].(int) + 1

	log.Println(result)

	return nil

}
