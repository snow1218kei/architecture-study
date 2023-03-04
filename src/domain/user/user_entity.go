package user

import (
	"fmt"
	"unicode/utf8"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared"
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

func NewUser(userMap map[string]interface{}) (*User, error) {
	if err := checkNameLength(userMap["name"].(string)); err != nil {
		return nil, err
	}

	if err := checkProfileLength(userMap["profile"].(string)); err != nil {
		return nil, err
	}

	user := &User{
		userID:    userMap["userID"].(UserID),
		name:      userMap["name"].(string),
		email:     userMap["email"].(shared.Email),
		password:  userMap["password"].(Password),
		profile:   userMap["profile"].(string),
		careers:   userMap["careers"].([]*Career),
		skills:    userMap["skills"].([]*Skill),
		createdAt: userMap["createdAt"].(shared.CreatedAt),
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
