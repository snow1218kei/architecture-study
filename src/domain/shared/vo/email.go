package shared

import (
	"fmt"
	"regexp"
)

const (
	emailPolicy = "^[a-zA-Z0-9_+-]+(\\.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\\.)+[a-zA-Z]{2,}$"
)

type Email string

func NewEmail(email string) (Email, error) {
	if err := isValidEmail(email); err != nil {
		return "", err
	}
	return Email(email), nil
}

func isValidEmail(email string) error {
	rxEmail := regexp.MustCompile(emailPolicy)
	if !rxEmail.MatchString(email) {
		return fmt.Errorf("無効なEmailアドレスです")
	}
	return nil
}
