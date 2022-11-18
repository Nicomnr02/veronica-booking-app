package service

import (
	"pos/pkg/domain/entity"
)

type CreateSizeRequest struct {
	Size string `json:"size"`
}

type UpdateSizeRequest struct {
	ID   int    `json:"id"`
	Size string `json:"size"`
}

type DeleteSizeRequest struct {
	ID int `json:"id"`
}

type SizeResponses struct {
	Sizes []entity.Sizes `json:"sizes"`
	Total int            `json:"total"`
}

type SizeResponse struct {
	Size entity.Size `json:"size"`
}
