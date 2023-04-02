package user

import (
	"fmt"
)

func CheckUserNameExistence(name string, userRepo UserRepository) error {
	_, err := userRepo.FindByName(name)

	if err != nil {
		return fmt.Errorf("既に存在するユーザ名です")
	}
	return nil
}
