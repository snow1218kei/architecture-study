package tag

import "context"

type Repository interface {
	FindByID(context.Context, TagID) (*Tag, error)
}
