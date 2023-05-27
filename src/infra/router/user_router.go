package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/controller"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/database"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/presenter"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase/userinput"
)

func NewUserRouter(g *gin.Engine) {
	g.POST("/users", func(ctx *gin.Context) {
		var in userinput.CreateUserInput
		if err := ctx.ShouldBindJSON(&in); err != nil {
			ctx.JSON(400, gin.H{"status": "bad request"})
			return
		}

		db, err := sqlx.Open("postgresql", "root/sample")
		if err != nil {
			ctx.Error(err)
			return
		}

		userRepoImpl := database.NewRdbUserRepository(db)
		if err := controller.NewUserController(
			presenter.NewUserPresenter(ctx), userRepoImpl).CreateUser(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}

	})
}
