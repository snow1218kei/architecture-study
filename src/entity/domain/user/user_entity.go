import (
	"errors"
)

type User struct {
  id UserID
  name string
	email Email
  password Password
	skills []Skill
	profile string
	career Career
}

func (user User) CheckNameLength error {
	if  user.name >= 255 {
		return errors.New("名前は255文字以下である必要があります")
  }
  return nil
}

func (user User) CheckProfileLength error {
	if  user.profile >= 2000 {
		return errors.New("プロフィールは2000文字以下である必要があります")
  }
  return nil
}
