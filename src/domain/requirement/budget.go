package requirement

import (
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

const minimumBudget uint16 = 1000

type Budget struct {
	budgetID   BudgetID
	lowerBound uint16
	upperBound uint16
}

type BudgetParams struct {
	LowerBound uint16
	UpperBound uint16
}

func newBudget(params *BudgetParams) (*Budget, error) {
	budget := &Budget{
		budgetID:   newBudgetID(),
		lowerBound: params.LowerBound,
		upperBound: params.UpperBound,
	}

	if err := budget.isValid(); err != nil {
		return nil, err
	}

	return budget, nil
}

func (b *Budget) isValid() error {
	if err := b.isAtLeastMinimum(); err != nil {
		return err
	}

	if err := b.isLowerBoundLessThanUpperBound(); err != nil {
		return err
	}

	return nil
}

func (b *Budget) isAtLeastMinimum() error {
	if b.lowerBound < minimumBudget || b.upperBound < minimumBudget {
		return apperr.BadRequestf("lowerBoundとupperBoundは%d円以上である必要があります: lowerBound=%d, upperBound=%d", minimumBudget, b.lowerBound, b.upperBound)
	}

	return nil
}

func (b *Budget) isLowerBoundLessThanUpperBound() error {
	if b.lowerBound > b.upperBound {
		return apperr.BadRequestf("lowerBoundはupperBound以下である必要があります: lowerBound=%d, upperBound=%d", b.lowerBound, b.upperBound)
	}

	return nil
}
