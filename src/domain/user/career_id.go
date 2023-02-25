package user

import (
	"github.com/google/uuid"
)

type CareerID string

func NewCareerID() CareerID {
	return CareerID(uuid.New().String())
}

func (careerId CareerID) String() string {
	return string(careerId)
}

func (careerId1 CareerID) Equal(careerId2 CareerID) bool {
	return careerId1 == careerId2
}
