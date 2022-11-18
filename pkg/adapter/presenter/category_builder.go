package presenter

import (
	"pos/pkg/domain/entity"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type CategoryBuilder struct{}

func (*CategoryBuilder) GetResponseCategories(s interface{}, totalData int) (interface{}, error) {
	var categories []entity.Categories
	err := mapstructure.Decode(s, &categories)
	if err != nil {
		panic(err)
	}

	return service.CategoryResponses{
		Categories: categories,
		Total:      totalData,
	}, nil
}

func (*CategoryBuilder) GetResponseCategory(u interface{}) (interface{}, error) {
	var category entity.Category
	err := mapstructure.Decode(u, &category)
	if err != nil {
		return nil, err
	}

	return service.CategoryResponse{
		Category: category,
	}, nil
}
