package service

import (
	"pos/pkg/domain/entity"
)

type CreateStockRequest struct {
	ProductID int    `json:"product_id"`
	Detail    string `json:"detail"`
	Quantity  int    `json:"quantity"`
}

type UpdateStockRequest struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Detail    string `json:"detail"`
	Quantity  int    `json:"quantity"`
}

type DeleteStockRequest struct {
	ID int `json:"id"`
}

type StockResponses struct {
	Stocks []entity.Stock `json:"stocks"`
	Total  int            `json:"total"`
}

type StockResponse struct {
	Stock entity.Stock `json:"stock"`
}
