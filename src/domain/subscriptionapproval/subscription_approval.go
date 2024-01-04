package subscriptionapproval

import (
	"github.com/yuuki-tsujimura/architecture-study/src/domain/mentoringplan"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"unicode/utf8"
)

const (
	ApprovalMessageMaxLength = 500
)

type SubscriptionApproval struct {
	subscriptionApprovalID SubscriptionApprovalID
	planID                 mentoringplan.MentoringPlanID
	message                string
}

type SubscriptionApprovalParams struct {
	PlanID  mentoringplan.MentoringPlanID
	Message string
}

func NewSubscriptionApproval(params SubscriptionApprovalParams) (*SubscriptionApproval, error) {
	if utf8.RuneCountInString(params.Message) > ApprovalMessageMaxLength {
		return nil, apperr.BadRequestf("Messageが%d文字を超えています: %d", ApprovalMessageMaxLength, params.Message)
	}

	return &SubscriptionApproval{
		subscriptionApprovalID: newSubscriptionApprovalID(),
		message:                params.Message,
	}, nil
}
