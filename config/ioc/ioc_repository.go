package ioc

import (
	"ApiGoLang/domain/entities"
	"ApiGoLang/domain/repository"
	"ApiGoLang/infra/db"
	rp "ApiGoLang/infra/repository"
	"github.com/golobby/container/v3"
)

func InitializeRepository() {
	var err error

	err = container.Transient(func() repository.IClienteRepository {
		return rp.NewClienteRepository(&rp.RepositoryBase[entities.Cliente]{DB: db.GetDb()}, db.GetDb())
	})

	if err != nil {
		panic(err)
	}
}
