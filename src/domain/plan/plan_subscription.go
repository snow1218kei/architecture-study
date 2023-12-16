package plan

type PlanSubscription struct {
	planSubscriptionID   PlanSubscriptionID
	subscriptionRequest  SubscriptionRequest
	subscriptionApproval *SubscriptionApproval
}

func newPlanSubscription(request SubscriptionRequest) (*PlanSubscription, error) {
	planSubscriptionID := newPlanSubscriptionID()

	planSubscription := &PlanSubscription{
		planSubscriptionID:   planSubscriptionID,
		subscriptionRequest:  request,
		subscriptionApproval: nil, // Not set during initialization
	}

	return planSubscription, nil
}

func (ps *PlanSubscription) setSubscriptionApproval(approval *SubscriptionApproval) {
	ps.subscriptionApproval = approval
}
