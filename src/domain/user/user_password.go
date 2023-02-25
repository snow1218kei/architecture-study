package user

import (
	"fmt"
	"regexp"
)

var (
	letterLengthPolicy = ".{12,}"
	alphabetPolicy = "[a-zA-Z]"
	digitPolicy = "[0-9]{6,}"
)

type Password string

func NewPassword(password string) (Password, error) {
	if err := validatePasswordLength(password); err != nil {
		return "", err
	}
	if err := validatePasswordContainsAlphabet(password); err != nil {
		return "", err
	}
	if err := validatePasswordContainsDigit(password); err != nil {
		return "", err
	}
	return Password(password), nil
}

func validatePasswordLength(password string) error {
	matched, _ := regexp.MatchString(letterLengthPolicy, password)
	if !matched {
			return fmt.Errorf("文字数は最低12文字以上でなければなりません")
	}
	return nil
}

func validatePasswordContainsAlphabet(password string) error {
	matched, _ := regexp.MatchString(alphabetPolicy, password)
	if !matched {
			return fmt.Errorf("英字が最低1文字は含まれていなければなりません")
	}
	return nil
}

func validatePasswordContainsDigit(password string) error {
	matched, _ := regexp.MatchString(digitPolicy, password)
	if !matched {
			return fmt.Errorf("数字が最低1文字は含まれていなければなりません")
	}
	return nil
}
