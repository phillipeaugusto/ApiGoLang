package services

import (
	"ApiGoLang/core/results"
	"ApiGoLang/domain/commands"
	_const "ApiGoLang/domain/const"
	"ApiGoLang/domain/dto/output"
	"ApiGoLang/domain/entities"
	"ApiGoLang/domain/repository"
	"github.com/golobby/container/v3"
)

type LoginService struct {
	rp repository.IClienteRepository
}

func NewLoginService() *LoginService {
	x := new(LoginService)
	_ = container.Resolve(&x.rp)
	return x
}

func (o *LoginService) Login(action *commands.LoginCommand) results.Result {

	var listErr, validData = action.Validate()
	if !validData {
		return results.Result{Success: false, Message: _const.DadoInvalido, Data: listErr}
	}

	obj := entities.NewClienteLogin(action.LoginInputDto.Email, action.LoginInputDto.Senha, true)

	err, objTemp := o.rp.GetClienteEmailSenhaLogin(obj.Email, obj.Senha)
	if err != nil {
		return results.Result{Success: false, Message: err.Error(), Data: nil}
	}

	if objTemp == nil {
		return results.Result{Success: false, Message: _const.DadosDeLoginInvalido, Data: nil}
	}

	return results.Result{Success: true, Message: _const.EntidadeRegistradaComSucesso, Data: output.DadosLoginOutPutDto{ClienteId: objTemp.ID, Token: "implementar_qualquer_jwt"}}
}
