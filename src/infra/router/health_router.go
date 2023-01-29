package router

import (
	"github.com/yuuki-tsujimura/architecture-study/src/interface/controller"
	"github.com/julienschmidt/httprouter"
)

func setHealthRouter(router *httprouter.Router) {
	router.GET("/health", controller.CheckHelthController)
}
