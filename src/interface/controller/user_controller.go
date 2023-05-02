package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	repoimpl "github.com/yuuki-tsujimura/architecture-study/src/infra/db/user"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/userusecase/userinput"
)

func CreateUserController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	var input userinput.CreateUserInput
	err = json.Unmarshal([]byte(body), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sqlx.Open("postgresql", "root/sample")
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	repository := repoimpl.NewRdbUserRepository(db)
	createUserUsecase := userusecase.NewCreateUserUseCase(repository)
	err = createUserUsecase.Exec(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
