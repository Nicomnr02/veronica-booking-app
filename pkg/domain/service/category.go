package service

import (
	"pos/pkg/domain/entity"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateCategoryRequest struct {
	Category string `json:"category"`
}

type UpdateCategoryRequest struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

type DeleteCategoryRequest struct {
	ID int `json:"id"`
}

type CategoryResponses struct {
	Categories []entity.Categories `json:"categories"`
	Total      int                 `json:"total"`
}

type CategoryResponse struct {
	Category entity.Category `json:"category"`
}

func (uc UpdateCategoryRequest) Validate() error {
	validate := validation.ValidateStruct(&uc, validation.Field(&uc.ID, validation.Required), validation.Field(&uc.Category, validation.Required))
	return validate
}
