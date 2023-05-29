package router

import (
	"ApiGoLang/api/v1/routers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var rt *mux.Router

func Initialize() {
	rt = mux.NewRouter().StrictSlash(true)
	rt.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	routers.Cliente(rt)
	routers.Login(rt)
}

func GetRouter() *mux.Router {
	return rt
}
