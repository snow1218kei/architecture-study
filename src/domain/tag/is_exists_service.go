package tag

import (
	"context"
)

type IsExistsService struct {
	repo TagRepository
}

func NewTagIDExistsService(repo TagRepository) *IsExistsService {
	return &IsExistsService{
		repo: repo,
	}
}

func (s *IsExistsService) Exec(ctx context.Context, tagIDs []TagID) (bool, []TagID, error) {
	tags, err := s.repo.FindByIDs(ctx, tagIDs)
	if err != nil {
		return false, nil, err
	}

	existedTagIDs := make([]TagID, len(tags))
	for _, tag := range tags {
		existedTagIDs = append(tagIDs, tag.tagID)
	}

	return len(tags) == len(tagIDs), existedTagIDs, nil
}
