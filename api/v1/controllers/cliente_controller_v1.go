package controllers

import (
	"ApiGoLang/application/services"
	"ApiGoLang/core/results"
	"ApiGoLang/core/utils"
	"ApiGoLang/domain/commands"
	"ApiGoLang/domain/commands/shared"
	_const "ApiGoLang/domain/const"
	"ApiGoLang/domain/dto/input"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type ClienteController struct{}

func NewClienteController() *ClienteController {
	return &ClienteController{}
}

// @Tags Cliente
// @Accept  json
// @Produce  json
// @Param request body input.ClienteInputDto true "request"
// @Success 200 {object} output.ClienteOutPutDto
// @Router /cliente [post]
func (o *ClienteController) Criar(w http.ResponseWriter, r *http.Request) {
	err, dto := utils.JsonForObject[input.ClienteInputDto](utils.ReaderForByte(r.Body))
	if err != nil {
		utils.ResultRequest(w, http.StatusBadRequest, results.Result{Success: false, Message: err.Error(), Data: nil})
		return
	}

	cmd := commands.NewCriarClienteCommand(dto)
	srv := services.NewClienteService().Create(cmd)

	if !srv.Success {
		utils.ResultRequest(w, http.StatusBadRequest, srv)
		return
	}

	utils.ResultRequest(w, http.StatusOK, srv)
}

// @Tags Cliente
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} output.ClienteOutPutDto
// @Router /cliente/{id} [get]
func (o *ClienteController) GetById(w http.ResponseWriter, r *http.Request) {

	var id, err = uuid.Parse(mux.Vars(r)["id"])
	if err.Error != nil {
		utils.ResultRequest(w, http.StatusInternalServerError, results.Result{Success: false, Message: _const.Erro, Data: err.Error()})
	}

	cmd := shared.NewBuscarViaIdCommand(id)
	srv := services.NewClienteService().GetById(cmd)

	if !srv.Success {
		utils.ResultRequest(w, http.StatusBadRequest, srv)
		return
	}

	utils.ResultRequest(w, http.StatusOK, srv)
}

// @Tags Cliente
// @Accept  json
// @Produce  json
// @Success 200 {object} []output.ClienteOutPutDto
// @Router /cliente [get]
func (o *ClienteController) GetAll(w http.ResponseWriter, r *http.Request) {

	srv := services.NewClienteService().GetAll()
	if !srv.Success {
		utils.ResultRequest(w, http.StatusBadRequest, srv)
		return
	}

	utils.ResultRequest(w, http.StatusOK, srv)
}
