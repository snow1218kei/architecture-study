package controller

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/presenter"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/planusecase"
)

type PlanController struct {
	presenter    presenter.PlanPresenter
	queryService planusecase.PlanQueryService
}

func NewPlanController(presenter presenter.PlanPresenter, queryService planusecase.PlanQueryService) *PlanController {
	return &PlanController{
		presenter,
		queryService,
	}
}

func (c PlanController) Index(ctx context.Context) error {
	usecase := planusecase.NewGetAllPlansUsecase(c.queryService)
	plan, err := usecase.Exec(ctx)
	if err != nil {
		return err
	}

	c.presenter.PlanList(plan)
	return nil
}
