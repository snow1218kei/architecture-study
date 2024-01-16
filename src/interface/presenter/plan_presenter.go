package presenter

import (
	"net/http"

	"github.com/yuuki-tsujimura/architecture-study/src/usecase/planusecase"
)

type PlanPresent struct {
	presenter Presenter
}

func NewPlanPresenter(presenter Presenter) PlanPresenter {
	return &PlanPresent{
		presenter,
	}
}

type PlanPresenter interface {
	PlanList(out []*planusecase.GetPlanDTO)
}

func (p *PlanPresent) PlanList(out []*planusecase.GetPlanDTO) {
	p.presenter.JSON(http.StatusOK, out)
}
