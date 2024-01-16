package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/controller"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/database"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/presenter"
)

func NewPlanRouter(g *gin.Engine) {
	g.GET("/plans", func(ctx *gin.Context) {
		db, err := sqlx.Open("postgresql", "root/sample")
		if err != nil {
			ctx.Error(err)
			return
		}

		if err := controller.NewPlanController(
			presenter.NewPlanPresenter(presenter.NewGinPresenter(ctx)),
			database.NewPlanQueryServiceImpl(db)).Index(ctx); err != nil {
			ctx.Error(err)
			return
		}
	})
}
