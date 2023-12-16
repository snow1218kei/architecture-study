package plan

import "context"

type PlanRepository interface {
	Store(context.Context, *MentoringPlan) error
	FindByID(context.Context, MentoringPlanID) (*MentoringPlan, error)
}
