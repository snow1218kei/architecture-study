package planusecase

import "context"

type PlanQueryService interface {
	GetAll(context.Context) ([]*GetPlanDTO, error)
}
