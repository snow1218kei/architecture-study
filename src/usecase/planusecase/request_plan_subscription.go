package planusecase

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/plan"
)

type RequestPlanSubscriptionUsecase struct {
	planRepo plan.PlanRepository
}

func NewRequestPlanSubscriptionUsecase(planRepo plan.PlanRepository) *RequestPlanSubscriptionUsecase {
	return &RequestPlanSubscriptionUsecase{
		planRepo,
	}
}

func (u *RequestPlanSubscriptionUsecase) Exec(ctx context.Context, input *SubscriptionInput) error {
	mentoringPlanID, err := plan.NewMentoringPlanIDByVal(input.mentoringPlanID)
	if err != nil {
		return err
	}

	mentoringPlan, err := u.planRepo.FindByID(ctx, mentoringPlanID)
	if err != nil {
		return err
	}

	mentoringPlan.AddSubscriptionRequest(input.message)

	if err := u.planRepo.Store(ctx, mentoringPlan); err != nil {
		return err
	}

	return nil
}
