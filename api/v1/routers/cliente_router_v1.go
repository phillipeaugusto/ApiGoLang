package routers

import (
	"ApiGoLang/api/v1/controllers"
	"github.com/gorilla/mux"
)

func Cliente(router *mux.Router) {
	r := router.PathPrefix("/v1/cliente").Subrouter()
	controller := controllers.NewClienteController()

	r.HandleFunc("", controller.Criar).Methods("POST")
	r.HandleFunc("/{id}", controller.GetById).Methods("GET")
	r.HandleFunc("/", controller.GetAll).Methods("GET")
}
