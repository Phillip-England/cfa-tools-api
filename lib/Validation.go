package lib

import (
	"fmt"
	"regexp"
)

func IsValidEmail(email string) (err error) {
	if email == "" {
		return fmt.Errorf("Email is required")
	}
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(email) {
		return fmt.Errorf("Please enter a valid email")
	}
	return nil
}

func IsValidPassword(password string) (err error) {
	if password == "" {
		return fmt.Errorf("Password is required")
	}
	if len(password) < 8 {
		return fmt.Errorf("Passwords must contain 8 or more character")
	}
	if len(password) > 64 {
		return fmt.Errorf("Passwords cannot contain more than 64 characters")
	}
	return nil
}