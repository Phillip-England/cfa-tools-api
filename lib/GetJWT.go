package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Claims jwt.RegisteredClaims
}

func GetJWT(data string) (signedString string, err error) {

	key := []byte(os.Getenv("JWT_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"iss":  "cfa-tools",
		"data": data,
	})

	signedString, err = token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedString, nil

}
