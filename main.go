package main

import (
	"ApiGoLang/config"
	"ApiGoLang/config/database"
	"ApiGoLang/config/database/cfg"
	"ApiGoLang/config/ioc"
	"ApiGoLang/config/router"
	"ApiGoLang/docs"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "APIGoLang"
	docs.SwaggerInfo.Description = "Api Exercitando Golang"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	config.LoadCfgGeneral()
	database.Connect(cfg.GetConfigDb())
	database.InitializeMigrations()
	ioc.InitializeRepository()
	router.Initialize()

	port := config.GetGeneralConfig().PortApi
	log.Printf("LISTEN ON PORT: %v", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%v", port), router.GetRouter())
}
