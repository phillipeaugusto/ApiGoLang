package repository

import (
	"ApiGoLang/core/repository/contracts"
	"ApiGoLang/domain/entities"
	"github.com/google/uuid"
)

type IClienteRepository interface {
	contracts.IRepository[entities.Cliente, entities.Cliente]
	GetClienteEmailSenhaLogin(email string, senha string) (error, *entities.Cliente)
	AtualizarStatus(id uuid.UUID, ativar bool) error
}
