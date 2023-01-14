package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CheckHelth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ヘルスチェックに成功しました。")
}

func main() {
	router := httprouter.New()
	router.GET("/health", CheckHelth)

	server := http.Server{
		Addr: "0.0.0.0:80",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("失敗しました: %v", err)
	}
}
