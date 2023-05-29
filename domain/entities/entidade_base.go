package entities

import (
	"ApiGoLang/core/domain/entities"
	"github.com/google/uuid"
	"time"
)

type EntidadeBase[TSource, TDestiny, TOutPut any] struct {
	entities.EntityAction[TSource, TDestiny, TOutPut]
	ID             uuid.UUID `gorm:"type:char(36);primary_key"`
	DataCriacao    time.Time `gorm:"not null"`
	DatalAlteracao time.Time `gorm:"not null"`
	Status         string    `gorm:"type:char(1);not null"`
}

func (e *EntidadeBase[TSource, TDestiny, TOutPut]) SetDefaultValues() *EntidadeBase[TSource, TDestiny, TOutPut] {
	e.ID = uuid.New()
	e.DataCriacao = time.Now()
	e.DatalAlteracao = time.Now()
	e.Status = "A"
	return e
}
