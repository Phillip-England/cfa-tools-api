package net

import (
	"errors"
)

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
