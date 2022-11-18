package category

import (
	"pos/pkg/domain/entity"
	"pos/pkg/domain/repository"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type CategoryUseCase struct {
	repo repository.CategoryRepository
	out  CategoryOutputPort
}

func NewCategoryUseCase(r repository.CategoryRepository, o CategoryOutputPort) *CategoryUseCase {
	return &CategoryUseCase{
		repo: r,
		out:  o,
	}
}

func (ct *CategoryUseCase) CreateCategory(req service.CreateCategoryRequest) error {
	var category entity.Category
	err := mapstructure.Decode(req, &category)
	if err != nil {
		return err
	}

	err = ct.repo.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (ct *CategoryUseCase) GetCategory(page, limit int) (interface{}, error) {
	data, err := ct.repo.GetCategory(page, limit)
	if err != nil {
		return nil, err
	}

	totalCategories, err := ct.repo.CountCategory()
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseCategories(data, totalCategories)
}

func (ct *CategoryUseCase) GetCategoryByID(ID int) (interface{}, error) {
	data, err := ct.repo.GetCategoryByID(ID)
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseCategory(data)
}

func (ct *CategoryUseCase) UpdateCategory(req service.UpdateCategoryRequest) error {
	var category entity.Category
	err := mapstructure.Decode(req, &category)
	if err != nil {
		return nil
	}

	err = ct.repo.UpdateCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func (ct *CategoryUseCase) DeleteCategory(req service.DeleteCategoryRequest) error {
	var category entity.Category
	err := mapstructure.Decode(req, &category)
	if err != nil {
		return err
	}

	err = ct.repo.DeleteCategory(category)
	if err != nil {
		return nil
	}

	return nil
}
