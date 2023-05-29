package repository

import (
	"ApiGoLang/core/repository/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math"
)

type RepositoryBase[TSource any] struct {
	DB gorm.DB
}

func (o *RepositoryBase[TSource]) Create(entity TSource) (error, bool) {
	result := o.DB.Create(entity)
	if result.Error != nil {
		return result.Error, false
	}
	return nil, true
}

func (o *RepositoryBase[TSource]) Update(entity TSource) (error, bool) {
	result := o.DB.Save(entity)
	if result.Error != nil {
		return result.Error, false
	}
	return nil, true
}

func (o *RepositoryBase[TSource]) Delete(entity TSource) (error, bool) {
	result := o.DB.Delete(entity)
	if result.Error != nil {
		return result.Error, false
	}
	return nil, true
}

func (o *RepositoryBase[TSource]) GetAll() (error, *[]TSource) {
	var objList []TSource
	result := o.DB.Find(&objList)
	if result.Error != nil {
		return result.Error, nil
	}

	return nil, &objList
}

func (o *RepositoryBase[TSource]) GetAllPagination(params pagination.PaginationParams) (error, *pagination.Pagination) {
	var objPagination pagination.Pagination
	var objList []TSource
	var obj TSource

	objPagination.Page = params.Page
	objPagination.Limit = params.Limit

	result := o.DB.Scopes(paginate(obj, objPagination, &o.DB)).Find(&objList)
	objPagination.Data = objList
	if result.Error != nil {
		return result.Error, nil
	}

	return nil, &objPagination
}
func (o *RepositoryBase[TSource]) GetById(id uuid.UUID) (error, *TSource) {
	var obj TSource
	result := o.DB.Find(&obj, "id = ?", id)
	if result.Error != nil {
		return result.Error, nil
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return nil, &obj
}

func paginate(value any, pagination pagination.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
