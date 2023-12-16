package plan

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type MentoringPlanID string

func newMentoringPlanID() MentoringPlanID {
	return MentoringPlanID(uuid.New().String())
}

func NewMentoringPlanIDByVal(val string) (MentoringPlanID, error) {
	if val == "" {
		return MentoringPlanID(""), apperr.BadRequest("mentorID must not be empty")
	}
	return MentoringPlanID(val), nil
}

func (planID MentoringPlanID) String() string {
	return string(planID)
}

func (planID1 MentoringPlanID) Equal(planID2 MentoringPlanID) bool {
	return planID1 == planID2
}
