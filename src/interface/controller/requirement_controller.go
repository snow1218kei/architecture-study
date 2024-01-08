package controller

import (
	"context"

	"github.com/yuuki-tsujimura/architecture-study/src/interface/presenter"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/requirementusecase"
)

type RequirementController struct {
	presenter    presenter.RequirementPresenter
	queryService requirementusecase.RequirementQueryService
}

func NewRequirementController(presenter presenter.RequirementPresenter, queryService requirementusecase.RequirementQueryService) *RequirementController {
	return &RequirementController{
		presenter,
		queryService,
	}
}

func (c RequirementController) Index(ctx context.Context) error {
	usecase := requirementusecase.NewGetAllRequirementsUsecase(c.queryService)
	requirementList, err := usecase.Exec(ctx)
	if err != nil {
		return err
	}

	c.presenter.RequirementList(requirementList)
	return nil
}
