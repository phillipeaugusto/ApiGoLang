package services

import (
	"ApiGoLang/core/results"
	"ApiGoLang/domain/commands"
	"ApiGoLang/domain/commands/shared"
	_const "ApiGoLang/domain/const"
	"ApiGoLang/domain/dto/output"
	"ApiGoLang/domain/entities"
	"ApiGoLang/domain/repository"
	"github.com/golobby/container/v3"
)

type ClienteService struct {
	rp repository.IClienteRepository
}

func NewClienteService() *ClienteService {
	x := new(ClienteService)
	_ = container.Resolve(&x.rp)
	return x
}

func (o *ClienteService) GetById(action *shared.BuscarViaIdCommand) results.Result {

	var listErr, validData = action.Validate()
	if !validData {
		return results.Result{Success: false, Message: _const.DadoJaExistente, Data: listErr}
	}

	err, obj := o.rp.GetById(action.Id)
	if err != nil {
		return results.Result{Success: false, Message: err.Error(), Data: nil}
	}

	if obj == nil {
		return results.Result{Success: false, Message: _const.EntidadeInexistente, Data: nil}
	}

	return results.Result{Success: true, Message: "", Data: obj.ConvertToObjectOutPut(*obj)}
}

func (o *ClienteService) GetAll() results.Result {

	err, objTemp := o.rp.GetAll()

	if err != nil {
		return results.Result{Success: false, Message: err.Error(), Data: nil}
	}

	var dataConvert = new([]output.ClienteOutPutDto)
	for _, item := range *objTemp {
		*dataConvert = append(*dataConvert, output.ClienteOutPutDto{Id: item.ID, Nome: item.Nome, Email: item.Email, Cpf: item.Cpf, Celular: item.Celular, DataAlteracao: item.DatalAlteracao, DataCriacao: item.DataCriacao, Status: item.Status})
	}

	return results.Result{Success: true, Message: "", Data: dataConvert}
}

func (o *ClienteService) Create(action *commands.CriarClienteCommand) results.Result {
	var listErr, validData = action.Validate()
	if !validData {
		return results.Result{Success: false, Message: _const.DadoInvalido, Data: listErr}
	}

	obj := entities.NewCliente(action.ClienteInputDto.Nome, action.ClienteInputDto.Email, action.ClienteInputDto.Cpf, action.ClienteInputDto.Celular, action.ClienteInputDto.Senha, true)

	err, objTemp := o.rp.Exists(*obj)
	if err != nil {
		return results.Result{Success: false, Message: err.Error(), Data: nil}
	}

	if objTemp != nil {
		return results.Result{Success: false, Message: _const.DadoJaExistente, Data: nil}
	}

	err, _ = o.rp.Create(*obj)
	if err != nil {
		return results.Result{Success: false, Message: err.Error(), Data: nil}
	}

	return results.Result{Success: true, Message: _const.EntidadeRegistradaComSucesso, Data: obj.ConvertToObject(*obj)}
}

func (o *ClienteService) AtualizarStatusViaId(action *shared.AtualizarStatusViaIdCommand) results.Result {
	var listErr, validData = action.Validate()

	if !validData {
		return results.Result{Success: false, Message: _const.DadoInvalido, Data: listErr}
	}

	err, objTemp := o.rp.GetById(action.Id)
	if err != nil {
		return results.Result{Success: false, Message: err.Error(), Data: nil}
	}

	if objTemp == nil {
		return results.Result{Success: false, Message: _const.EntidadeInexistente, Data: nil}
	}

	err = o.rp.AtualizarStatus(action.Id, action.Ativar)
	if err != nil {
		return results.Result{Success: false, Message: err.Error(), Data: nil}
	}

	return results.Result{Success: true, Message: _const.EntidadeAtualizadaComSucesso, Data: nil}
}
