package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CheckHelth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "ヘルスチェックに成功しました。")
}

func main() {
	router := httprouter.New()
	router.Get("/health", CheckHelth)
	log.Fatal(http.ListenAndServe(":8000", router))

	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// 	Handler: router,
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Fatalf("失敗しました: %v", err)
	// }
}
