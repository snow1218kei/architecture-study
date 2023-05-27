package user

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type CareerID string

func newCareerID() CareerID {
	return CareerID(uuid.New().String())
}

func NewCareerIDByVal(val string) (CareerID, error) {
	if val == "" {
		return CareerID(""), apperr.BadRequest("careerID must not be empty")
	}
	return CareerID(val), nil
}

func (careerId CareerID) String() string {
	return string(careerId)
}

func (careerId1 CareerID) Equal(careerId2 CareerID) bool {
	return careerId1 == careerId2
}
