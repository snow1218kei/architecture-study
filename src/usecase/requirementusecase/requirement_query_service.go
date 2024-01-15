package requirementusecase

import "context"

type RequirementQueryService interface {
	GetAll(context.Context) ([]*GetMentorRequirementDTO, error)
}
