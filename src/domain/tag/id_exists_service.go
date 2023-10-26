package tag

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type IDExistsService struct {
	repo Repository
}

func NewTagIDExistsService(repo Repository) *IDExistsService {
	return &IDExistsService{
		repo: repo,
	}
}

func (s *IDExistsService) Exec(ctx context.Context, tagID TagID) (bool, error) {
	tag, err := s.repo.FindByID(ctx, tagID)

	if err != nil {
		if apperr.Is[*apperr.NotFoundErr](err) {
			return false, nil
		}

		return false, err
	}

	return tag != nil, nil
}
