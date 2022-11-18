package product

import (
	"pos/pkg/domain/entity"
	"pos/pkg/domain/repository"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type ProductUseCase struct {
	repo repository.ProductRepository
	out  ProductOutputPort
}

func NewProductUseCase(r repository.ProductRepository, o ProductOutputPort) *ProductUseCase {
	return &ProductUseCase{
		repo: r,
		out:  o,
	}
}

func (ct *ProductUseCase) CreateProduct(req service.CreateProductRequest) error {
	var product entity.Product
	err := mapstructure.Decode(req, &product)
	if err != nil {
		return err
	}

	err = ct.repo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (ct *ProductUseCase) GetProduct(page, limit int) (interface{}, error) {
	data, err := ct.repo.GetProduct(page, limit)
	if err != nil {
		return nil, err
	}

	totalProducts, err := ct.repo.CountProduct()
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseProducts(data, totalProducts)
}

func (ct *ProductUseCase) GetProductByID(ID int) (interface{}, error) {
	data, err := ct.repo.GetProductByID(ID)
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseProduct(data)
}

func (ct *ProductUseCase) UpdateProduct(req service.UpdateProductRequest) error {
	var product entity.Product
	err := mapstructure.Decode(req, &product)
	if err != nil {
		return nil
	}

	err = ct.repo.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (ct *ProductUseCase) DeleteProduct(req service.DeleteProductRequest) error {
	var product entity.Product
	err := mapstructure.Decode(req, &product)
	if err != nil {
		return err
	}

	err = ct.repo.DeleteProduct(product)
	if err != nil {
		return nil
	}

	return nil
}
