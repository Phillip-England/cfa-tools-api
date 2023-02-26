package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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

func DecodeJWT(signedString string) (jwtData interface{}, err error) {
	keyFunc := func(signedString *jwt.Token) (interface{}, error) {
		key := []byte(os.Getenv("JWT_KEY"))
		return key, nil
	}
	token, err := jwt.Parse(signedString, keyFunc)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	jwtData, ok := claims["data"]
	if !ok {
		return nil, err
	}
	return jwtData, nil
}

func Decrypt(ciphertext []byte) (value []byte, err error) {
	var key []byte = []byte(os.Getenv("CRYPTO_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func Encrypt(value []byte) (ciphertext []byte, err error) {
	var key []byte = []byte(os.Getenv("CRYPTO_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext = make([]byte, aes.BlockSize+len(value))
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], value)
	return ciphertext, nil
}

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

func IsCSRF(token string) (err error) {
	randomRunes := []rune(token)
	err = errors.New("invalid csrf token")
	if string(randomRunes[43]) != "G" {
		return err
	}
	if string(randomRunes[55]) != "O" {
		return err
	}
	return nil
}