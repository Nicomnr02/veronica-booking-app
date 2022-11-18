package category

import (
	"pos/pkg/domain/service"
	"pos/pkg/usecase/category"

	"github.com/sirupsen/logrus"
)

type CategoryService struct {
	CategoryRepository category.CategoryInputPort
}

func NewCategoryService(ct category.CategoryInputPort) *CategoryService {
	return &CategoryService{
		CategoryRepository: ct,
	}
}
func (ct *CategoryService) CreateCategory(req service.CreateCategoryRequest) error {
	err := ct.CategoryRepository.CreateCategory(req)
	if err != nil {
		return err
	}

	return nil
}

func (ct *CategoryService) GetCategory(page, limit int) (interface{}, error) {
	category, err := ct.CategoryRepository.GetCategory(page, limit)
	if err != nil {
		logrus.Error("[GetCategory] error get Category in Controller")
		return nil, err
	}
	return category, err
}

func (ct *CategoryService) GetCategoryByID(ID int) (interface{}, error) {
	category, err := ct.CategoryRepository.GetCategoryByID(ID)
	if err != nil {
		logrus.Error("[GetCategoryByID] error get category by id in Controller")
		return nil, err
	}
	return category, nil
}

func (ct *CategoryService) DeleteCategory(req service.DeleteCategoryRequest) error {
	err := ct.CategoryRepository.DeleteCategory(req)
	if err != nil {
		return nil
	}
	return err
}

func (ct *CategoryService) UpdateCategory(req service.UpdateCategoryRequest) error {
	err := ct.CategoryRepository.UpdateCategory(req)
	if err != nil {
		return err
	}

	return nil
}
