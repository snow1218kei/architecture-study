package router

import (
	"go-module/src/interface/controller"
	"github.com/julienschmidt/httprouter"
)

func GetHealth() *httprouter.Router {
	router := httprouter.New()
	router.GET("/health", controller.CheckHelth)

	return router
}
