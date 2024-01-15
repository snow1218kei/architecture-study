package presenter

import (
	"net/http"

	"github.com/yuuki-tsujimura/architecture-study/src/usecase/requirementusecase"
)

type RequirementPresent struct {
	presenter Presenter
}

func NewRequirementPresenter(presenter Presenter) RequirementPresenter {
	return &RequirementPresent{
		presenter,
	}
}

type RequirementPresenter interface {
	RequirementList(out []*requirementusecase.GetMentorRequirementDTO)
}

func (p *RequirementPresent) RequirementList(out []*requirementusecase.GetMentorRequirementDTO) {
	p.presenter.JSON(http.StatusOK, out)
}
