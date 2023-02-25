package user

import (
	"fmt"
)

type User struct {
	userID   UserID
	name     string
	email    Email
	password Password
	profile  string
	careers   []*Career
	skills   []*Skill
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
		userID:   userMap["userID"].(UserID),
		name:     userMap["name"].(string),
		email:    userMap["email"].(Email),
		password: userMap["password"].(Password),
		profile:  userMap["profile"].(string),
		careers:   userMap["careers"].([]*Career),
		skills:   userMap["skills"].([]*Skill),
	}
	return user, nil
}

func checkNameLength(name string) error {
	if len(name) > 255 {
		return fmt.Errorf("名前は255文字以下である必要があります。(現在%d文字)", len(name))
	}
	return nil
}

func checkProfileLength(profile string) error {
	if len(profile) >= 2000 {
		return fmt.Errorf("プロフィールは2000文字以下である必要があります。(現在%d文字)", len(profile))
	}
	return nil
}
