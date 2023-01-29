package main

import (
	"net/http"
	"log"

	"github.com/yuuki-tsujimura/architecture-study/src/infra/router"
)

func main() {
	router := router.Router()

	server := http.Server{
		Addr: "0.0.0.0:80",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("失敗しました: %v", err)
	}
}
