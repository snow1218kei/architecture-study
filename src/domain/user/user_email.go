package user

import (
	"fmt"
	"regexp"
)

type Email string

func NewEmail(email string) (Email, error) {
	e := Email(email)

	if !e.IsValidEmail() {
		return e, fmt.Errorf("有効なメールアドレスではありません")
	}
	return e, nil
}

func (email Email) IsValidEmail() bool {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*.)+[a-zA-Z]{2,}$")
	return rxEmail.MatchString(string(email))
}
