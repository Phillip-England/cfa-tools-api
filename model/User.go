package model

import (
	"strings"
	"time"

	"github.com/phillip-england/go-http/lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
}

type UserResponse struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email"`
}

func BuildUser(email string, password string) (user User, err error) {
	email = strings.ToLower(email)
	if err != nil {
		return User{}, err
	}
	user = User{
		Email:    email,
		Password: password,
	}
	user.Timestamp()
	return user, nil
}

func (v *User) Timestamp() {
	now := time.Now()
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	v.UpdatedAt = now
}

func (v *User) EncryptPassword() (err error) {
	encryptedPassword, err := lib.Encrypt([]byte(v.Password))
	if err != nil {
		return err
	}
	v.Password = string(encryptedPassword)
	return nil
} 

func (v *User) GetDecryptedPassword() (password string, err error) {
	passwordBytes, err := lib.Decrypt([]byte(v.Password))
	if err != nil {
		return "", err
	}
	password = string(passwordBytes)
	return password, nil
}

func (v *User) Validate() (err error) {
	err = lib.IsValidEmail(v.Email)
	if err != nil {
		return err
	}
	err = lib.IsValidPassword(v.Password)
	if err != nil {
		return err
	}
	return nil
}
