package requirement

import "context"

type ReqRepository interface {
	Store(context.Context, *MentorRequirement) error
}
