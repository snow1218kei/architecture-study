package subscriptionapproval

import "context"

type SubscriptionApprovalRepository interface {
	Store(context.Context, *SubscriptionApproval) error
}
