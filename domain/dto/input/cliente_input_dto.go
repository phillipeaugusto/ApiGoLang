package input

import (
	"ApiGoLang/core/domain/entities"
	o "ApiGoLang/domain/entities"
)

type ClienteInputDto struct {
	entities.EntityAction[ClienteInputDto, o.Cliente, o.Cliente]
	Nome    string
	Email   string
	Cpf     string
	Celular string
	Senha   string
}

func NewClienteInputDto(nome string, email string, cpf string, celular string, senha string) *ClienteInputDto {
	return &ClienteInputDto{Nome: nome, Email: email, Cpf: cpf, Celular: celular, Senha: senha}
}
