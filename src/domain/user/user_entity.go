package user

import (
	"fmt"
	"time"
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

type UserData struct {
	UserID    string
	Name      string
	Email     string
	Password  string
	Profile   string
	Careers   []*CareerData
	Skills    []*SkillData
	CreatedAt time.Time
}

func newUser(userInput UserInput) (*User, error) {
	if err := checkNameLength(userInput.Name); err != nil {
		return nil, err
	}

	if err := checkProfileLength(userInput.Profile); err != nil {
		return nil, err
	}

	return &User{
		userID:    userInput.UserID,
		name:      userInput.Name,
		email:     userInput.Email,
		password:  userInput.Password,
		profile:   userInput.Profile,
		careers:   userInput.Careers,
		skills:    userInput.Skills,
		createdAt: userInput.CreatedAt,
	}, nil
}

func ReconstructUserFromData(userData UserData) (*User, error) {
	return &User{
		userID:    UserID(userData.UserID),
		name:      userData.Name,
		email:     shared.Email(userData.Email),
		password:  Password(userData.Password),
		profile:   userData.Profile,
		careers:   ReconstructCareersFromData(userData.Careers),
		skills:    ReconstructSkillsFromData(userData.Skills),
		createdAt: shared.CreatedAt(userData.CreatedAt),
	}, nil
}

func ConvertUserToUserData(user *User) UserData {
	return UserData{
		UserID:    user.userID.String(),
		Name:      user.name,
		Email:     string(user.email),
		Password:  string(user.password),
		Profile:   user.profile,
		Careers:   ConvertCareersToCareerData(user.careers),
		Skills:    ConvertSkillsToSkillData(user.skills),
		CreatedAt: user.createdAt.Value(),
	}
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
