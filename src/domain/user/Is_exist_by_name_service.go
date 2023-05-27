package user

import (
	"context"

	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type IsExistByNameService struct {
	repo UserRepository
}

func NewIsExistByNameService(repo UserRepository) *IsExistByNameService {
	return &IsExistByNameService{
		repo: repo,
	}
}

func (ds *IsExistByNameService) Exec(ctx context.Context, name string) (bool, error) {
	user, err := ds.repo.FindByName(ctx, name)

	if err != nil {
		if apperr.Is(err, &apperr.NotFoundErr{}) {
			return false, nil
		}
		return false, err
	}
	return user != nil, nil
}
