package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cares struct {
	ID                    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User                  primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Location              primitive.ObjectID `json:"location,omitempty" bson:"location,omitempty"`
	CreatedAt             time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt             time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	GuestName             string             `json:"guestName" bson:"guestName"`
	OrderNumber           string             `json:"orderNumber" bson:"orderNumber"`
	Incident              string             `json:"incident" bson:"incident"`
	ReplacementAction     string             `json:"replacementAction" bson:"replacementAction"`
	ReplacementCode       int                `json:"replacementCode,omitempty" bson:"replacementCode,omitempty"`
	ReplacementCodeString string             `json:"replacementCodeString,omitempty" bson:"replacementCodeString,omitempty"`
}

func (v *Cares) Timestamp() {
	now := time.Now()
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	v.UpdatedAt = now
}

func (v *Cares) SetReplacementCode(ctx context.Context, coll *mongo.Collection) (err error) {

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "user", Value: v.User}}}},
		{{Key: "$sort", Value: bson.D{{Key: "replacementCode", Value: -1}}}},
		{{Key: "$limit", Value: 1}},
		{{Key: "$project", Value: bson.D{{Key: "replacementCode", Value: 1}, {Key: "_id", Value: nil}}}},
	}

	var result []Cares
	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if cursor.All(ctx, &result); err != nil {
		return err
	}

	if len(result) == 0 {
		v.ReplacementCode = 1000
		v.ReplacementCodeString = fmt.Sprintf("%s%d", "CFA", v.ReplacementCode)
		return nil
	}

	v.ReplacementCode = result[0].ReplacementCode + 1
	v.ReplacementCodeString = fmt.Sprintf("%s%d", "CFA", v.ReplacementCode)

	return nil

}
