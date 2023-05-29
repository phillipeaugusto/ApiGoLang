package contracts

import (
	"ApiGoLang/core/results"
)

type IService[T any] interface {
	Service(action T) results.Result
}
