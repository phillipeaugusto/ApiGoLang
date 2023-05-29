package database

import (
	"ApiGoLang/config/database/cfg"
	"ApiGoLang/domain/entities"
	"ApiGoLang/infra/db"
)

func Connect(cfg cfg.DbConfig) {
	err := db.OpenConnection(cfg)
	if err != nil {
		panic(err)
	}
}

func InitializeMigrations() {
	x := db.GetDb()

	err := x.AutoMigrate(
		new(entities.Cliente),
	)

	if err != nil {
		panic(err)
	}

}
