package repository

import "pos/pkg/domain/entity"

type CategoryRepository interface {
	CreateCategory(req entity.Category) error
	GetCategory(page, limt int) (interface{}, error)
	GetCategoryByID(ID int) (interface{}, error)
	UpdateCategory(req entity.Category) error
	DeleteCategory(req entity.Category) error
	CountCategory() (total int, err error)
}
