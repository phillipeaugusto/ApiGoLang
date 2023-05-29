package input

import (
	"ApiGoLang/core/domain/entities"
	entities2 "ApiGoLang/domain/entities"
)

type LoginInputDto struct {
	entities.EntityAction[LoginInputDto, LoginInputDto, entities2.Cliente]
	Email string
	Senha string
}

func NewLoginInputDto(email string, senha string) *LoginInputDto {
	return &LoginInputDto{Email: email, Senha: senha}
}
