package shared

import (
	"ApiGoLang/core/validations"
	"github.com/google/uuid"
)

type AtualizarStatusViaIdCommand struct {
	Id     uuid.UUID
	Ativar bool
}

func NewAtualizarStatusViaIdCommand(id uuid.UUID, ativar bool) *AtualizarStatusViaIdCommand {
	return &AtualizarStatusViaIdCommand{Id: id, Ativar: ativar}
}

func (obj *AtualizarStatusViaIdCommand) Validate() ([]validations.ErrorData, bool) {
	var data = new([]validations.ValidationData)

	*data = append(*data,
		validations.ValidationData{Valid: obj.Id.String() == "", Field: "Id", Message: "Inválido, verifique"},
	)

	return validations.Notifications(data), validations.Validate(data)
}
