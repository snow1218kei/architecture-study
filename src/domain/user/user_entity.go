package user

import (
	"fmt"
	"unicode/utf8"

	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
)

type User struct {
	userID    UserID
	name      string
	email     shared.Email
	password  Password
	profile   string
	careers   []*Career
	skills    []*Skill
	createdAt shared.CreatedAt
}

type UserParams struct {
	Name     string
	Email    string
	Password string
	Profile  string
}

type UserInput struct {
	UserID    UserID
	Name      string
	Email     shared.Email
	Password  Password
	Profile   string
	Careers   []*Career
	Skills    []*Skill
	CreatedAt shared.CreatedAt
}

func newUser(userInput UserInput) (*User, error) {
	if err := checkNameLength(userInput.Name); err != nil {
		return nil, err
	}

	if err := checkProfileLength(userInput.Profile); err != nil {
		return nil, err
	}

	user := &User{
		userID:    userInput.UserID,
		name:      userInput.Name,
		email:     userInput.Email,
		password:  userInput.Password,
		profile:   userInput.Profile,
		careers:   userInput.Careers,
		skills:    userInput.Skills,
		createdAt: userInput.CreatedAt,
	}
	return user, nil
}

func checkNameLength(name string) error {
	if utf8.RuneCountInString(name) > 255 {
		return fmt.Errorf("名前は255文字以下である必要があります。(現在%d文字)", utf8.RuneCountInString(name))
	}
	return nil
}

func checkProfileLength(profile string) error {
	if utf8.RuneCountInString(profile) >= 2000 {
		return fmt.Errorf("プロフィールは2000文字以下である必要があります。(現在%d文字)", utf8.RuneCountInString(profile))
	}
	return nil
}
