package user

import (
	"fmt"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/id"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase"
)

type User struct {
	UserId   id.UserId
	Name     string
	Password Password
	Profile  string
	Career   Career
	Skills   []Skill
}

func NewUser(input usecase.CreateUserInput) User {
	userId := id.NewUserId()
	name, _ := CheckNameLength(input.User.Name)
	password, _ := NewPassword(input.User.Password)
	profile, _ := CheckProfileLength(input.User.Profile)
	career := *NewCareer(input.Career)
	skills := PrepareSkills(input)

	return User{
		UserId:   userId,
		Name:     name,
		Password: password,
		Profile:  profile,
		Career:   career,
		Skills:   skills,
	}
}

func CheckNameLength(name string) (string, error) {
	if len(name) > 255 {
		return "", fmt.Errorf("名前は255文字以下である必要があります。(現在%d文字)", len(name))
	}
	return name, nil
}

func CheckProfileLength(profile string) (string, error) {
	if len(profile) >= 2000 {
		return "", fmt.Errorf("プロフィールは2000文字以下である必要があります。(現在%d文字)", len(profile))
	}
	return profile, nil
}

func PrepareSkills(input usecase.CreateUserInput) []Skill {
	var skills []Skill

	for _, skill := range input.Skills {
		skills = append(skills, *NewSkill(skill))
	}
	return skills
}
