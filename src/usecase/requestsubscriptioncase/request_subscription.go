package requestsubscriptioncase

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/mentoringplan"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/subscriptionrequest"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type RequestPlanSubscriptionUsecase struct {
	planRepo    mentoringplan.PlanRepository
	requestRepo subscriptionrequest.SubscriptionRequestRepository
}

func NewRequestPlanSubscriptionUsecase(planRepo mentoringplan.PlanRepository, requestRepo subscriptionrequest.SubscriptionRequestRepository) *RequestPlanSubscriptionUsecase {
	return &RequestPlanSubscriptionUsecase{
		planRepo,
		requestRepo,
	}
}

func (u *RequestPlanSubscriptionUsecase) Exec(ctx context.Context, input *RequestSubscriptionInput) error {
	mentoringPlanID, err := mentoringplan.NewMentoringPlanIDByVal(input.mentoringPlanID)
	if err != nil {
		return err
	}

	planExistsService := mentoringplan.NewIsExistByIDService(u.planRepo)
	isPlanExist, err := planExistsService.Run(ctx, mentoringPlanID)
	if err != nil {
		return err
	}
	if !isPlanExist {
		return apperr.BadRequestWrapf(err, "このmentoringPlanは存在しません: %s", input.mentoringPlanID)
	}

	params := subscriptionrequest.SubscriptionRequestParams{
		PlanID:  mentoringPlanID,
		Message: input.message,
	}

	subscriptionRequest, err := subscriptionrequest.NewSubscriptionRequest(params)
	if err != nil {
		return err
	}

	if err := u.requestRepo.Store(ctx, subscriptionRequest); err != nil {
		return err
	}

	return nil
}
