package planusecase

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/plan"
)

type ApprovePlanSubscriptionUsecase struct {
	planRepo plan.PlanRepository
}

func NewApprovePlanSubscriptionUsecase(planRepo plan.PlanRepository) *ApprovePlanSubscriptionUsecase {
	return &ApprovePlanSubscriptionUsecase{
		planRepo,
	}
}

func (u *ApprovePlanSubscriptionUsecase) Exec(ctx context.Context, input *SubscriptionInput) error {
	mentoringPlanID, err := plan.NewMentoringPlanIDByVal(input.mentoringPlanID)
	if err != nil {
		return err
	}

	mentoringPlan, err := u.planRepo.FindByID(ctx, mentoringPlanID)
	if err != nil {
		return err
	}

	mentoringPlan.AddSubscriptionApproval(input.message)

	if err := u.planRepo.Store(ctx, mentoringPlan); err != nil {
		return err
	}

	return nil
}
