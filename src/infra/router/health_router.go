package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yuuki-tsujimura/architecture-study/src/interface/controller"
)

func setHealthRouter(router *httprouter.Router) {
	router.GET("/health", controller.CheckHelthController)
}
