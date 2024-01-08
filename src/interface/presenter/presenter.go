package presenter

import "github.com/gin-gonic/gin"

type Presenter interface {
	JSON(code int, obj any)
}

type ginPresenter struct {
	ctx *gin.Context
}

func NewGinPresenter(ctx *gin.Context) Presenter {
	return &ginPresenter{ctx: ctx}
}

func (p *ginPresenter) JSON(code int, obj any) {
	p.ctx.JSON(code, obj)
}
