package plan

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type PlanSubscriptionID string

func newPlanSubscriptionID() PlanSubscriptionID {
	return PlanSubscriptionID(uuid.New().String())
}

func NewPlanSubscriptionIDByVal(val string) (PlanSubscriptionID, error) {
	if val == "" {
		return PlanSubscriptionID(""), apperr.BadRequest("planSubscriptionID must not be empty")
	}
	return PlanSubscriptionID(val), nil
}

func (planSubscriptionID PlanSubscriptionID) String() string {
	return string(planSubscriptionID)
}

func (planSubscriptionID1 PlanSubscriptionID) Equal(planSubscriptionID2 PlanSubscriptionID) bool {
	return planSubscriptionID1 == planSubscriptionID2
}
