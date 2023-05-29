package contracts

import (
	"ApiGoLang/core/repository/pagination"
	"github.com/google/uuid"
)

type IRepositoryBase[TEntity any] interface {
	Create(entity TEntity) (error, bool)
	Update(entity TEntity) (error, bool)
	Delete(entity TEntity) (error, bool)
	GetAll() (error, *[]TEntity)
	GetAllPagination(params pagination.PaginationParams) (error, *pagination.Pagination)
	GetById(id uuid.UUID) (error, *TEntity)
}
