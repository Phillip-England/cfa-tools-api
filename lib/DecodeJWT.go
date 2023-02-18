package lib

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func DecodeJWT(signedString string) (jwtData interface{}, err error) {
	
	keyFunc := func(signedString *jwt.Token) (interface{}, error) {
		key := []byte(os.Getenv("JWT_KEY"))
		return key, nil
	}

	token, err := jwt.Parse(signedString, keyFunc)
	if (err != nil) {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	jwtData, ok := claims["data"]
	if (!ok) {
		return nil, err
	}
	
	return jwtData, nil


}