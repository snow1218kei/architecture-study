package presenter

import (
	"net/http"

	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase/useroutput"
)

type userPresent struct {
	delivery Presenter
}

func NewUserPresenter(p Presenter) UserPresenter {
	return &userPresent{
		delivery: p,
	}
}

type UserPresenter interface {
	Create(out *useroutput.CreateUserOutput)
}

func (p *userPresent) Create(out *useroutput.CreateUserOutput) {
	p.delivery.JSON(http.StatusCreated, out)
}
