package requirement

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type MentorID string

func newMentorID() MentorID {
	return MentorID(uuid.New().String())
}

func NewMentorIDByVal(val string) (MentorID, error) {
	if val == "" {
		return MentorID(""), apperr.BadRequest("mentorID must not be empty")
	}
	return MentorID(val), nil
}

func (mentorID MentorID) String() string {
	return string(mentorID)
}

func (mentorID1 MentorID) Equal(mentorID2 MentorID) bool {
	return mentorID1 == mentorID2
}
