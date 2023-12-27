package approvesubscriptionusecase

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/subscriptionapproval"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/subscriptionrequest"
)

type ApprovePlanSubscriptionUsecase struct {
	requestRepo  subscriptionrequest.SubscriptionRequestRepository
	approvalRepo subscriptionapproval.SubscriptionApprovalRepository
}

func NewApprovePlanSubscriptionUsecase(requestRepo subscriptionrequest.SubscriptionRequestRepository, approvalRepo subscriptionapproval.SubscriptionApprovalRepository) *ApprovePlanSubscriptionUsecase {
	return &ApprovePlanSubscriptionUsecase{
		requestRepo,
		approvalRepo,
	}
}

func (u *ApprovePlanSubscriptionUsecase) Exec(ctx context.Context, input *ApproveSubscriptionInput) error {
	requestID, err := subscriptionrequest.NewSubscriptionRequestIDByVal(input.subscriptionRequestID)
	if err != nil {
		return err
	}

	subscriptionRequest, err := u.requestRepo.FindByID(ctx, requestID)
	if err != nil {
		return err
	}

	params := subscriptionapproval.SubscriptionApprovalParams{
		PlanID:  subscriptionRequest.GetPlanID(),
		Message: input.message,
	}

	subscriptionApproval, err := subscriptionapproval.NewSubscriptionApproval(params)
	if err != nil {
		return err
	}

	if err := u.approvalRepo.Store(ctx, subscriptionApproval); err != nil {
		return err
	}

	return nil
}
