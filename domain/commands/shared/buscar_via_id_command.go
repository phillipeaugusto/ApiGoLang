package shared

import (
	"ApiGoLang/core/validations"
	"github.com/google/uuid"
)

type BuscarViaIdCommand struct {
	Id uuid.UUID
}

func NewBuscarViaIdCommand(id uuid.UUID) *BuscarViaIdCommand {
	return &BuscarViaIdCommand{Id: id}
}

func (obj *BuscarViaIdCommand) Validate() ([]validations.ErrorData, bool) {
	var data = new([]validations.ValidationData)

	*data = append(*data,
		validations.ValidationData{Valid: obj.Id.String() == "", Field: "Id", Message: "Inválido, verifique"},
	)

	return validations.Notifications(data), validations.Validate(data)
}
