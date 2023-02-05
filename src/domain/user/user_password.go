package user

import (
	"fmt"
	"regexp"
)

type Password string

func NewPassword(password string) (Password, error) {
	p := Password(password)

	if !p.IsCompliantPasswordPolicy() {
		return p, fmt.Errorf("パスワードは規約に適合していません")
	}
	return p, nil
}

func (password Password) IsCompliantPasswordPolicy() bool {
	re := regexp.MustCompile("^(?=.*[A-Za-z])(?=.*[0-9])[A-Za-z0-9]{6,}$")
	return re.MatchString(string(password))
}
