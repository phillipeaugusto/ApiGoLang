package dto

import "ApiGoLang/core/domain/entities"

type DtoBase[TSource, TDestiny, TOutPut any] struct {
	entities.EntityAction[TSource, TDestiny, TDestiny]
}
