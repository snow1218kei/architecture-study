package plan

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type SubscriptionApproval struct {
	subscriptionApprovalID SubscriptionApprovalID
	message                string
}

func newSubscriptionApproval(message string) (*SubscriptionApproval, error) {
	if len(message) > MaxLength {
		return nil, apperr.BadRequestf("Messageが500文字を超えています: %d", message)
	}

	return &SubscriptionApproval{
		subscriptionApprovalID: newSubscriptionApprovalID(),
		message:                message,
	}, nil
}
