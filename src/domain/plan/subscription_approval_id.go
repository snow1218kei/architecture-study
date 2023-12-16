package plan

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type SubscriptionApprovalID string

func newSubscriptionApprovalID() SubscriptionApprovalID {
	return SubscriptionApprovalID(uuid.New().String())
}

func NewSubscriptionApprovalIDByVal(val string) (SubscriptionApprovalID, error) {
	if val == "" {
		return SubscriptionApprovalID(""), apperr.BadRequest("subscriptionApprovalID must not be empty")
	}
	return SubscriptionApprovalID(val), nil
}

func (approvalID SubscriptionApprovalID) String() string {
	return string(approvalID)
}

func (approvalID1 SubscriptionApprovalID) Equal(approvalID2 SubscriptionApprovalID) bool {
	return approvalID1 == approvalID2
}
