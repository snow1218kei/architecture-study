package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CheckHelthController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ヘルスチェックに成功しました。")
}
