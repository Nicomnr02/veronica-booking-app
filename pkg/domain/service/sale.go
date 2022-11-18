package service

import (
	"pos/pkg/shared/util"
)

type Sale struct {
	ID         int                `json:"id"`
	TotalPrice int                `json:"total_price"`
	Discount   int                `json:"discount"`
	FinalPrice int                `json:"final_price"`
	Cash       int                `json:"cash"`
	Remaining  int                `json:"remaining"`
	Note       string             `json:"note"`
	UserID     int                `json:"user_id"`
	Items      []util.PropertyMap `json:"items"`
	CreatedAt  string             `json:"created_at"`
	UpdatedAt  string             `json:"updated_at"`
}

type CreateSaleRequest struct {
	TotalPrice int                `json:"total_price"`
	Discount   int                `json:"discount"`
	FinalPrice int                `json:"final_price"`
	Cash       int                `json:"cash"`
	Remaining  int                `json:"remaining"`
	Note       string             `json:"note"`
	UserID     int                `json:"user_id"`
	Items      []util.PropertyMap `json:"items"`
}

type UpdateSaleRequest struct {
	ID         int                `json:"id"`
	TotalPrice int                `json:"total_price"`
	Discount   int                `json:"discount"`
	FinalPrice int                `json:"final_price"`
	Cash       int                `json:"cash"`
	Remaining  int                `json:"remaining"`
	Note       string             `json:"note"`
	UserID     int                `json:"user_id"`
	Items      []util.PropertyMap `json:"items"`
}

type DeleteSaleRequest struct {
	ID int `json:"id"`
}

type SaleResponses struct {
	Sales []Sale `json:"Sales"`
	Total int    `json:"total"`
}

type SaleResponse struct {
	Sale Sale `json:"Sale"`
}
