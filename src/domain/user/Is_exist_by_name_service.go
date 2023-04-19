package user

import (
	"context"
	"errors"
)

var IsNotFound = errors.New("not found")

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
		if errors.Is(err, IsNotFound) {
			return false, nil
		}
		return false, err
	}
	return user != nil, nil
}

