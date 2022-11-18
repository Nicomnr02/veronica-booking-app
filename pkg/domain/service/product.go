package service

import (
	"pos/pkg/domain/entity"
)

type CreateProductRequest struct {
	CategoryID    int    `json:"category_id"`
	SizeID        int    `json:"size_id"`
	Barcode       string `json:"Barcode"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	DetailProduct string `json:"detail_product"`
}

type UpdateProductRequest struct {
	ID            int    `json:"id"`
	CategoryID    int    `json:"category_id"`
	SizeID        int    `json:"size_id"`
	Barcode       string `json:"Barcode"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	DetailProduct string `json:"detail_product"`
}

type DeleteProductRequest struct {
	ID int `json:"id"`
}

type ProductResponses struct {
	Products []entity.Products `json:"products"`
	Total    int               `json:"total"`
}

type ProductResponse struct {
	Product entity.Product `json:"product"`
}

// func (uc UpdateProductRequest) Validate() error {
// 	validate := validation.ValidateStruct(&uc, validation.Field(&uc.ID, validation.Required), validation.Field(&uc.Product, validation.Required))
// 	return validate
// }
