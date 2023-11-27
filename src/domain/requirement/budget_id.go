package requirement

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type BudgetID string

func newBudgetID() BudgetID {
	return BudgetID(uuid.New().String())
}

func NewBudgetIDByVal(val string) (BudgetID, error) {
	if val == "" {
		return BudgetID(""), apperr.BadRequest("BudgetID must not be empty")
	}
	return BudgetID(val), nil
}

func (budgetID BudgetID) String() string {
	return string(budgetID)
}

func (budgetID1 BudgetID) Equal(budgetID2 BudgetID) bool {
	return budgetID1 == budgetID2
}
