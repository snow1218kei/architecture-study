package requirement

import "context"

type Repository interface {
	Store(context.Context, *MentorRequirement) error
}
