package model

import (
	"strings"
	"time"

	"github.com/phillip-england/go-http/lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string `json:"email" bson:"email"`
}

func BuildUser(email string, password string) (user User, err error) {
	encryptedPassword, err := lib.Encrypt([]byte(password))
	email = strings.ToLower(email)
	if err != nil {
		return User{}, err
	}
	user = User{
		Email: email,
		Password: string(encryptedPassword),
	}
	user.Timestamp()
	return user, nil
}

func (user *User) Timestamp() {
	now := time.Now()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	user.UpdatedAt = now
}

func (user *User) GetDecryptedPassword() (password string, err error) {
	passwordBytes, err := lib.Decrypt([]byte(user.Password))
	if err != nil {
		return "", err
	}
	password = string(passwordBytes)
	return password, nil
}
