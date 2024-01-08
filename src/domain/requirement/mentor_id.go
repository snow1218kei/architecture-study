package requirement

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type MentorRequirementID string

func newMentorID() MentorRequirementID {
	return MentorRequirementID(uuid.New().String())
}

func NewMentorRequirementIDByVal(val string) (MentorRequirementID, error) {
	if val == "" {
		return MentorRequirementID(""), apperr.BadRequest("mentorID must not be empty")
	}
	return MentorRequirementID(val), nil
}

func (mentorRequirementID MentorRequirementID) String() string {
	return string(mentorRequirementID)
}

func (mentorRequirementID1 MentorRequirementID) Equal(mentorRequirementID2 MentorRequirementID) bool {
	return mentorRequirementID1 == mentorRequirementID2
}
