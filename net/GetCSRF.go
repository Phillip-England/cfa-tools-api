package net

import (
	"crypto/rand"
	"encoding/base64"
)


func GetCSRF() (token string, err error) {

	length := 64

	randomBytes := make([]byte, length)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	randomRunes := []rune(randomString)

	randomRunes[43] = 'G'
	randomRunes[55] = 'O'

	token = string(randomRunes)

	return token, nil
}