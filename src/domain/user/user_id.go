package user

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type UserID string

func newUserID() UserID {
	return UserID(uuid.New().String())
}

func NewUserIDByVal(val string) (UserID, error) {
	if val == "" {
		return UserID(""), apperr.BadRequest("userID must not be empty")
	}
	return UserID(val), nil
}

func (userId UserID) String() string {
	return string(userId)
}

func (userId1 UserID) Equal(userId2 UserID) bool {
	return userId1 == userId2
}
