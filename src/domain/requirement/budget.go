package requirement

import "github.com/yuuki-tsujimura/architecture-study/src/support/apperr"

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
	if b.lowerBound < 1000 || b.upperBound < 1000 {
		return apperr.BadRequestf("lowerBoundとupperBoundは%1000以上である必要があります: %d", b.lowerBound, b.upperBound)
	}

	return nil
}

func (b *Budget) isLowerBoundLessThanUpperBound() error {
	if b.lowerBound > b.upperBound {
		return apperr.BadRequestf("lowerBoundはupperBound以下である必要があります: %d", b.lowerBound, b.upperBound)
	}

	return nil
}
