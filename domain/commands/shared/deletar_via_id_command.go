package shared

import (
	"ApiGoLang/core/validations"
	"github.com/google/uuid"
)

type DeleteViaIdCommand struct {
	Id uuid.UUID
}

func NewDeleteViaIdCommand(id uuid.UUID) *DeleteViaIdCommand {
	return &DeleteViaIdCommand{Id: id}
}

func (obj *DeleteViaIdCommand) Validate() ([]validations.ErrorData, bool) {
	var data = new([]validations.ValidationData)

	*data = append(*data,
		validations.ValidationData{Valid: obj.Id.String() == "", Field: "Id", Message: "Inválido, verifique"},
	)

	return validations.Notifications(data), validations.Validate(data)
}
