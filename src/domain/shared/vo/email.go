package shared

import (
	"regexp"

	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
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
		return apperr.BadRequestf("無効なEmailアドレスです: %s", email)
	}
	return nil
}
