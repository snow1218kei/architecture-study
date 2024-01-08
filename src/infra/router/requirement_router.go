package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/controller"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/database"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/presenter"
)

func NewRequirementRouter(g *gin.Engine) {
	g.GET("/requirements", func(ctx *gin.Context) {
		db, err := sqlx.Open("postgresql", "root/sample")
		if err != nil {
			ctx.Error(err)
			return
		}

		if err := controller.NewRequirementController(
			presenter.NewRequirementPresenter(presenter.NewGinPresenter(ctx)),
			database.NewRequirementQueryServiceImpl(db)).Index(ctx); err != nil {
			ctx.Error(err)
			return
		}
	})
}
