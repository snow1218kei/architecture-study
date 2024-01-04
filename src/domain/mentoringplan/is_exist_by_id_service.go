package mentoringplan

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type IsExistByIDService struct {
	repo PlanRepository
}

func NewIsExistByIDService(repo PlanRepository) *IsExistByIDService {
	return &IsExistByIDService{
		repo: repo,
	}
}

func (ds *IsExistByIDService) Run(ctx context.Context, planID MentoringPlanID) (bool, error) {
	user, err := ds.repo.FindByID(ctx, planID)

	if err != nil {
		if apperr.Is[*apperr.NotFoundErr](err) {
			return false, nil
		}
		return false, err
	}

	return user != nil, nil
}
