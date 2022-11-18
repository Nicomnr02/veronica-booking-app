package stock

import "pos/pkg/domain/service"

type StockInputPort interface {
	CreateStock(req service.CreateStockRequest) error
	GetStock(page, limit int) (interface{}, error)
	GetStockByID(ID int) (interface{}, error)
	UpdateStock(req service.UpdateStockRequest) error
	DeleteStock(req service.DeleteStockRequest) error
}

type StockOutputPort interface {
	GetResponseStocks(interface{}, int) (interface{}, error)
	GetResponseStock(interface{}) (interface{}, error)
}
