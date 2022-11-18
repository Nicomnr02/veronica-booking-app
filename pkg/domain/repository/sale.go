package repository

import "pos/pkg/domain/entity"

type SaleRepository interface {
	CreateSale(req entity.Sale) error
	GetSale(page, limt int) (interface{}, error)
	GetSaleByID(ID int) (interface{}, error)
	UpdateSale(req entity.Sale) error
	DeleteSale(req entity.Sale) error
	CountSale() (total int, err error)
}
