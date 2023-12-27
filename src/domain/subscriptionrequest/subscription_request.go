package subscriptionrequest

import (
	"github.com/yuuki-tsujimura/architecture-study/src/domain/mentoringplan"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"unicode/utf8"
)

const (
	RequestMessageMaxLength = 500
)

type SubscriptionRequest struct {
	subscriptionRequestID SubscriptionRequestID
	planID                mentoringplan.MentoringPlanID
	message               string
}

type SubscriptionRequestParams struct {
	PlanID  mentoringplan.MentoringPlanID
	Message string
}

func NewSubscriptionRequest(params SubscriptionRequestParams) (*SubscriptionRequest, error) {
	if utf8.RuneCountInString(params.Message) > RequestMessageMaxLength {
		return nil, apperr.BadRequestf("Messageが500文字を超えています: %s", params.Message)
	}

	return &SubscriptionRequest{
		subscriptionRequestID: newSubscriptionRequestID(),
		message:               params.Message,
	}, nil
}

func (s SubscriptionRequest) GetPlanID() mentoringplan.MentoringPlanID {
	return s.planID
}
