package commands

import (
	"ApiGoLang/core/validations"
	"ApiGoLang/domain/dto/input"
)

type LoginCommand struct {
	LoginInputDto input.LoginInputDto
}

func NewLoginCommand(loginInputDto input.LoginInputDto) *LoginCommand {
	return &LoginCommand{LoginInputDto: loginInputDto}
}

func (obj *LoginCommand) Validate() ([]validations.ErrorData, bool) {
	var data = new([]validations.ValidationData)

	*data = append(*data,
		validations.ValidationData{Valid: obj.LoginInputDto.Email == "", Field: "Email", Message: "Inválido, verifique."},
		validations.ValidationData{Valid: obj.LoginInputDto.Senha == "", Field: "Senha", Message: "Inválida, verifique."},
	)

	return validations.Notifications(data), validations.Validate(data)
}
