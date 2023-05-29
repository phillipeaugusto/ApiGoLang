package commands

import (
	"ApiGoLang/core/validations"
	"ApiGoLang/domain/dto/input"
)

type CriarClienteCommand struct {
	ClienteInputDto input.ClienteInputDto
}

func NewCriarClienteCommand(clienteInputDto input.ClienteInputDto) *CriarClienteCommand {
	return &CriarClienteCommand{ClienteInputDto: clienteInputDto}
}

func (obj *CriarClienteCommand) Validate() ([]validations.ErrorData, bool) {
	var data = new([]validations.ValidationData)

	*data = append(*data,
		validations.ValidationData{Valid: obj.ClienteInputDto.Cpf == "", Field: "Cpf", Message: "Inválido."},
		validations.ValidationData{Valid: len(obj.ClienteInputDto.Cpf) < 11, Field: "CpfDigitos", Message: "Quantidade de digitos, Inválido(s)."},
		validations.ValidationData{Valid: obj.ClienteInputDto.Nome == "", Field: "Nome", Message: "Inválido."},
		validations.ValidationData{Valid: len(obj.ClienteInputDto.Nome) < 3, Field: "NomeCaracteres", Message: "Nome Inválido, Verifique"},
		validations.ValidationData{Valid: obj.ClienteInputDto.Email == "", Field: "Email", Message: "Inválido"},
		validations.ValidationData{Valid: obj.ClienteInputDto.Celular == "", Field: "Celular", Message: "Inválido"},
		validations.ValidationData{Valid: len(obj.ClienteInputDto.Celular) < 12, Field: "CelularNúmero", Message: "Numero do celular está, Inválido"},
	)

	return validations.Notifications(data), validations.Validate(data)
}
