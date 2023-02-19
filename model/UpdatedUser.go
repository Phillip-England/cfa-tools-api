package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdatedUser struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Email     string `json:"email" bson:"email"`
	CurrentPassword string `json:"current_password" bson:"current_password"`
	NewPassword  string `json:"new_password" bson:"new_password"`
}

func (v *UpdatedUser) Timestamp() {
	now := time.Now()
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	v.UpdatedAt = now
}