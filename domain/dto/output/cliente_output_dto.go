package output

import (
	"ApiGoLang/core/domain/entities"
	e "ApiGoLang/domain/entities"
	"github.com/google/uuid"
	"time"
)

type ClienteOutPutDto struct {
	entities.EntityAction[ClienteOutPutDto, ClienteOutPutDto, e.Cliente]
	Id            uuid.UUID
	Status        string
	Nome          string
	Email         string
	Cpf           string
	Celular       string
	DataCriacao   time.Time
	DataAlteracao time.Time
}

func NewClienteOutPutdto(id uuid.UUID, status string, nome string, email string, cpf string, celular string, dataCriacao time.Time, dataAlteracao time.Time) *ClienteOutPutDto {
	return &ClienteOutPutDto{Id: id, Status: status, Nome: nome, Email: email, Cpf: cpf, Celular: celular, DataCriacao: dataCriacao, DataAlteracao: dataAlteracao}
}
