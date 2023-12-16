package plan

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type SubscriptionRequestID string

func newSubscriptionRequestID() SubscriptionRequestID {
	return SubscriptionRequestID(uuid.New().String())
}

func NewSubscriptionRequestIDByVal(val string) (SubscriptionRequestID, error) {
	if val == "" {
		return SubscriptionRequestID(""), apperr.BadRequest("SubscriptionRequestID must not be empty")
	}
	return SubscriptionRequestID(val), nil
}

func (requestID SubscriptionRequestID) String() string {
	return string(requestID)
}

func (requestID1 SubscriptionRequestID) Equal(requestID2 SubscriptionRequestID) bool {
	return requestID1 == requestID2
}
