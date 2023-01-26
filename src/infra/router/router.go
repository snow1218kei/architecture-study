package router

import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	setHealthRouter(router)

	return router
}
