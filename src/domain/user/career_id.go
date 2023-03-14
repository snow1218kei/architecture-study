package user

import (
	"errors"

	"github.com/google/uuid"
)

type CareerID string

func newCareerID() CareerID {
	return CareerID(uuid.New().String())
}

func NewCareerIDByVal(val string) (CareerID, error) {
	if val == "" {
		return CareerID(""), errors.New("careerID must not be empty")
	}
	return CareerID(val), nil
}

func (careerId CareerID) String() string {
	return string(careerId)
}

func (careerId1 CareerID) Equal(careerId2 CareerID) bool {
	return careerId1 == careerId2
}
