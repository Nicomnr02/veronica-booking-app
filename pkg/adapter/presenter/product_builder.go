package presenter

import (
	"pos/pkg/domain/entity"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type ProductBuilder struct{}

func (*ProductBuilder) GetResponseProducts(s interface{}, totalData int) (interface{}, error) {
	var products []entity.Products
	err := mapstructure.Decode(s, &products)
	if err != nil {
		panic(err)
	}

	return service.ProductResponses{
		Products: products,
		Total:    totalData,
	}, nil
}

func (*ProductBuilder) GetResponseProduct(u interface{}) (interface{}, error) {
	var product entity.Product
	err := mapstructure.Decode(u, &product)
	if err != nil {
		return nil, err
	}

	return service.ProductResponse{
		Product: product,
	}, nil
}
