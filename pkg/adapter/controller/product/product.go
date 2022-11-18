package product

import (
	"pos/pkg/domain/service"
	"pos/pkg/usecase/product"

	"github.com/sirupsen/logrus"
)

type ProductService struct {
	ProductRepository product.ProductInputPort
}

func NewProductService(ct product.ProductInputPort) *ProductService {
	return &ProductService{
		ProductRepository: ct,
	}
}
func (ct *ProductService) CreateProduct(req service.CreateProductRequest) error {
	err := ct.ProductRepository.CreateProduct(req)
	if err != nil {
		return err
	}

	return nil
}

func (ct *ProductService) GetProduct(page, limit int) (interface{}, error) {
	product, err := ct.ProductRepository.GetProduct(page, limit)
	if err != nil {
		logrus.Error("[GetProduct] error get Product in Controller")
		return nil, err
	}
	return product, err
}

func (ct *ProductService) GetProductByID(ID int) (interface{}, error) {
	product, err := ct.ProductRepository.GetProductByID(ID)
	if err != nil {
		logrus.Error("[GetProductByID] error get product by id in Controller")
		return nil, err
	}
	return product, nil
}

func (ct *ProductService) DeleteProduct(req service.DeleteProductRequest) error {
	err := ct.ProductRepository.DeleteProduct(req)
	if err != nil {
		return nil
	}
	return err
}

func (ct *ProductService) UpdateProduct(req service.UpdateProductRequest) error {
	err := ct.ProductRepository.UpdateProduct(req)
	if err != nil {
		return err
	}

	return nil
}
