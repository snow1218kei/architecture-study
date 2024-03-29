package user

import (
	"time"
	"unicode/utf8"

	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

const (
	maxNameLength    = 255
	maxProfileLength = 2000
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

type UserUpdateParams struct {
	Name     *string
	Email    *string
	Password *string
	Profile  *string
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

func (u *User) ID() UserID {
	return u.userID
}

func ReconstructUserFromData(userData UserData) *User {
	return &User{
		userID:    UserID(userData.UserID),
		name:      userData.Name,
		email:     shared.Email(userData.Email),
		password:  Password(userData.Password),
		profile:   userData.Profile,
		careers:   ReconstructCareersFromData(userData.Careers),
		skills:    ReconstructSkillsFromData(userData.Skills),
		createdAt: shared.CreatedAt(userData.CreatedAt),
	}
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
	if l := utf8.RuneCountInString(name); l > maxNameLength {
		return apperr.BadRequestf("名前は%d文字以下である必要があります。(現在%d文字)", maxNameLength, l)
	}
	return nil
}

func checkProfileLength(profile string) error {
	if l := utf8.RuneCountInString(profile); l >= maxProfileLength {
		return apperr.BadRequestf("プロフィールは%d文字以下である必要があります。(現在%d文字)", maxProfileLength, l)
	}
	return nil
}

func (u *User) update(params UserUpdateParams) error {
	if params.Name != nil {
		if err := checkNameLength(*params.Name); err != nil {
			return err
		}

		u.name = *params.Name
	}

	if params.Email != nil {
		email, err := shared.NewEmail(*params.Email)
		if err != nil {
			return err
		}

		u.email = email
	}

	if params.Password != nil {
		password, err := NewPassword(*params.Password)
		if err != nil {
			return err
		}

		u.password = password
	}

	if params.Profile != nil {
		if err := checkProfileLength(*params.Profile); err != nil {
			return err
		}

		u.profile = *params.Profile
	}

	return nil
}
