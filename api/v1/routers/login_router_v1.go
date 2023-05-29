package routers

import (
	"ApiGoLang/api/v1/controllers"
	"github.com/gorilla/mux"
)

func Login(router *mux.Router) {
	r := router.PathPrefix("/v1/login").Subrouter()
	controller := controllers.NewLoginController()
	r.HandleFunc("", controller.Login).Methods("POST")
}
