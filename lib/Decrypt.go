package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func Decrypt(ciphertext []byte) (value []byte, err error) {

	var key []byte = []byte(os.Getenv("CRYPTO_KEY"))

	block, err := aes.NewCipher(key)
	if (err != nil) {
		return nil, err
	}

	if (len(ciphertext) < aes.BlockSize) {
		return nil, fmt.Errorf("ciphertext is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil

}