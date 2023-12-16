package plan

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

type SubscriptionRequest struct {
	subscriptionRequestID SubscriptionRequestID
	message               string
}

func newSubscriptionRequest(message string) (*SubscriptionRequest, error) {
	if len(message) > MaxLength {
		return nil, apperr.BadRequestf("Messageが500文字を超えています: %d", message)
	}

	return &SubscriptionRequest{
		subscriptionRequestID: newSubscriptionRequestID(),
		message:               message,
	}, nil
}
