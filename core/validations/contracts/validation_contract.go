package contracts

import "ApiGoLang/core/validations"

type IValidate interface {
	Validate() ([]validations.ErrorData, bool)
}
