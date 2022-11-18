package category

import "pos/pkg/domain/service"

type CategoryInputPort interface {
	CreateCategory(req service.CreateCategoryRequest) error
	GetCategory(page, limit int) (interface{}, error)
	GetCategoryByID(ID int) (interface{}, error)
	UpdateCategory(req service.UpdateCategoryRequest) error
	DeleteCategory(req service.DeleteCategoryRequest) error
}

type CategoryOutputPort interface {
	GetResponseCategories(interface{}, int) (interface{}, error)
	GetResponseCategory(interface{}) (interface{}, error)
}
