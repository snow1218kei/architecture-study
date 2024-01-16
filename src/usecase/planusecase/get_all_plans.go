package planusecase

import "context"

type GetAllPlansUsecase struct {
	queryService PlanQueryService
}

func NewGetAllPlansUsecase(queryService PlanQueryService) *GetAllPlansUsecase {
	return &GetAllPlansUsecase{
		queryService,
	}
}

func (u *GetAllPlansUsecase) Exec(ctx context.Context) ([]*GetPlanDTO, error) {
	plans, err := u.queryService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return plans, nil
}
