package id

import (
	"github.com/google/uuid"
)

type UserId string

func NewUserId() UserId {
	uuid, _ := uuid.NewRandom()
	return UserId(uuid.String())
}

func (userId UserId) String() string {
	return string(userId)
}

func (userId1 UserId) Equal(userId2 UserId) bool {
	return userId1 == userId2
}
