package user

import (
	"context"
	"fmt"
)

func CheckUserNameExistence(ctx context.Context, name string, userRepo UserRepository) error {
	_, err := userRepo.FindByName(ctx, name)

	if err != nil {
		return fmt.Errorf("既に存在するユーザ名です")
	}
	return nil
}
