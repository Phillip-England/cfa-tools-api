package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func Encrypt(value []byte) (ciphertext []byte, err error) {

	// getting key
	var key []byte = []byte(os.Getenv("CRYPTO_KEY"))

	// creating a new cipher
	block, err := aes.NewCipher(key)
	if (err != nil) {
		return nil, err
	}

	// generating unique algorithm
	ciphertext = make([]byte, aes.BlockSize+len(value))
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if (err != nil) {
		return nil, err
	}

	// encrypting the value
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], value)
	return ciphertext, nil


}