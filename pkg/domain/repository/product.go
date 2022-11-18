package repository

import "pos/pkg/domain/entity"

type ProductRepository interface {
	CreateProduct(req entity.Product) error
	GetProduct(page, limt int) (interface{}, error)
	GetProductByID(ID int) (interface{}, error)
	UpdateProduct(req entity.Product) error
	UpdateProductStock(req entity.Product) error
	DeleteProduct(req entity.Product) error
	CountProduct() (total int, err error)
}
