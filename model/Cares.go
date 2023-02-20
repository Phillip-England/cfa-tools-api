package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ReplacementCode   string             `json:"replacementCode,omitempty" bson:"replacementCode,omitempty"`
}

func (v *Cares) Timestamp() {
	now := time.Now()
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	v.UpdatedAt = now
}
