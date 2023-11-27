package tag

import "context"

type TagRepository interface {
	FindByIDs(ctx context.Context, ids []TagID) ([]*Tag, error)
}
