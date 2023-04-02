package user

import (
	"fmt"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

type RdbUserRepository struct {
	users []*user.User
}

func NewRdbUserRepository() *RdbUserRepository {
	return &RdbUserRepository{
			users: make([]*user.User, 0),
	}
}

func (repo *RdbUserRepository) Store(user *user.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *RdbUserRepository) FindByName(name string) error {
	return fmt.Errorf("user not found")
}
