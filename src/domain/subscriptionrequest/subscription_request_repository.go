package subscriptionrequest

import "context"

type SubscriptionRequestRepository interface {
	Store(context.Context, *SubscriptionRequest) error
	FindByID(context.Context, SubscriptionRequestID) (*SubscriptionRequest, error)
}
