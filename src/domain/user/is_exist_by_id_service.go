package user

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type IsExistByIDService struct {
	repo UserRepository
}

func NewIsExistByIDService(repo UserRepository) *IsExistByIDService {
	return &IsExistByIDService{
		repo: repo,
	}
}

func (ds *IsExistByIDService) Run(ctx context.Context, userID UserID) (bool, error) {
	user, err := ds.repo.FindByID(ctx, userID)

	if err != nil {
		if apperr.Is[*apperr.NotFoundErr](err) {
			return false, nil
		}
		return false, err
	}

	return user != nil, nil
}
