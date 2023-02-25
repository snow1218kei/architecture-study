package user

import (
	"github.com/google/uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func (userId UserID) String() string {
	return string(userId)
}

func (userId1 UserID) Equal(userId2 UserID) bool {
	return userId1 == userId2
}
