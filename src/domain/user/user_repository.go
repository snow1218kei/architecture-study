package user

import (
	"context"
)

type UserRepository interface {
	Store(context.Context, *User) error
	FindByName(context.Context, string) (*User, error)
	FindByID(context.Context, UserID) (*User, error)
	Update(context.Context, *User) error
}
