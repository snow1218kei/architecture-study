package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yuuki-tsujimura/architecture-study/src/infra/db"
	"github.com/yuuki-tsujimura/architecture-study/src/infra/middleware"
	"github.com/yuuki-tsujimura/architecture-study/src/infra/router"
)

func main() {
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	g := gin.Default()

	g.Use(middleware.HandleErrorMiddleware())
	g.Use(middleware.DBMiddleware(conn))

	router.NewUserRouter(g)

	log.Fatal(g.Run(":8000"))
}
