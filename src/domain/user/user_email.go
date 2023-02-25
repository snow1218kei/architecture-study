package user

import (
	"fmt"
	"regexp"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared"
)

type Email shared.Email

func NewEmail(email string) (Email, error) {
	if err := IsValidEmail(email); err != nil {
		return "", err
	}
	return Email(email), nil
}

func IsValidEmail(email string) error {
	rxEmail := regexp.MustCompile(shared.EmailPolicy)
	if !rxEmail.MatchString(email) {
			return fmt.Errorf("invalid email address")
	}
	return nil
}
