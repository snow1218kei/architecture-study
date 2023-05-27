package controller

import (
	"context"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/presenter"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase/userinput"
)

type userController struct {
	delivery presenter.UserPresenter
	userRepo user.UserRepository
}

func NewUserController(p presenter.UserPresenter, userRepo user.UserRepository) *userController {
	return &userController{
		delivery: p,
		userRepo: userRepo,
	}
}

func (c *userController) CreateUser(ctx context.Context, in *userinput.CreateUserInput) error {
	usecase := userusecase.NewCreateUserUseCase(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.Create(out)
	return nil
}
