package repository

import "pos/pkg/domain/entity"

type StockRepository interface {
	CreateStock(req entity.Stock) error
	GetStock(page, limt int) (interface{}, error)
	GetStockByID(ID int) (interface{}, error)
	UpdateStock(req entity.Stock) error
	DeleteStock(req entity.Stock) error
	CountStock() (total int, err error)
}
