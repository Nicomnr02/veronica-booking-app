package sale

import "pos/pkg/domain/service"

type SaleInputPort interface {
	CreateSale(req service.CreateSaleRequest) error
	GetSale(page, limit int) (interface{}, error)
	GetSaleByID(ID int) (interface{}, error)
	UpdateSale(req service.UpdateSaleRequest) error
	DeleteSale(req service.DeleteSaleRequest) error
}

type SaleOutputPort interface {
	GetResponseSales(interface{}, int) (interface{}, error)
	GetResponseSale(interface{}) (interface{}, error)
}
