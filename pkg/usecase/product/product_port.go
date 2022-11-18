package product

import "pos/pkg/domain/service"

type ProductInputPort interface {
	CreateProduct(req service.CreateProductRequest) error
	GetProduct(page, limit int) (interface{}, error)
	GetProductByID(ID int) (interface{}, error)
	DeleteProduct(req service.DeleteProductRequest) error
	UpdateProduct(req service.UpdateProductRequest) error
}

type ProductOutputPort interface {
	GetResponseProducts(interface{}, int) (interface{}, error)
	GetResponseProduct(interface{}) (interface{}, error)
}
