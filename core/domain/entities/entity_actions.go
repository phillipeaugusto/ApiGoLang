package entities

import (
	"ApiGoLang/core/utils"
)

type EntityAction[TSource, TDestiny, TOutPut any] struct {
}

func (e *EntityAction[TSource, TDestiny, TOutPut]) ConvertToObject(o TSource) *TDestiny {
	err, obj := utils.JsonForObject[TDestiny](utils.ObjectForByte(o))
	if err != nil {
		return nil
	}
	return &obj

}

func (e *EntityAction[TSource, TDestiny, TOutPut]) ConvertToObjectOutPut(o TSource) *TOutPut {
	err, obj := utils.JsonForObject[TOutPut](utils.ObjectForByte(o))
	if err != nil {
		return nil
	}
	return &obj
}
