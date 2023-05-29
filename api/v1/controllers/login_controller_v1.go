package controllers

import (
	"ApiGoLang/application/services"
	"ApiGoLang/core/results"
	"ApiGoLang/core/utils"
	"ApiGoLang/domain/commands"
	"ApiGoLang/domain/dto/input"
	"net/http"
)

type LoginController struct {
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

// @Tags Login
// @Accept  json
// @Produce  json
// @Param request body input.LoginInputDto true "request"
// @Success 200 {object} output.DadosLoginOutPutDto
// @Router /login [post]
func (o *LoginController) Login(w http.ResponseWriter, r *http.Request) {

	err, dto := utils.JsonForObject[input.LoginInputDto](utils.ReaderForByte(r.Body))
	if err != nil {
		utils.ResultRequest(w, http.StatusBadRequest, results.Result{Success: false, Message: err.Error(), Data: nil})
		return
	}

	cmd := commands.NewLoginCommand(dto)
	srv := services.NewLoginService().Login(cmd)

	if !srv.Success {
		utils.ResultRequest(w, http.StatusBadRequest, srv)
		return
	}

	utils.ResultRequest(w, http.StatusOK, srv)
}
