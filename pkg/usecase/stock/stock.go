package stock

import (
	"pos/pkg/domain/entity"
	"pos/pkg/domain/repository"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type StockUseCase struct {
	repo repository.StockRepository
	out  StockOutputPort
}

func NewStockUseCase(r repository.StockRepository, o StockOutputPort) *StockUseCase {
	return &StockUseCase{
		repo: r,
		out:  o,
	}
}

func (ct *StockUseCase) CreateStock(req service.CreateStockRequest) error {
	var stock entity.Stock
	err := mapstructure.Decode(req, &stock)
	if err != nil {
		return err
	}

	err = ct.repo.CreateStock(stock)
	if err != nil {
		return err
	}
	return nil
}

func (ct *StockUseCase) GetStock(page, limit int) (interface{}, error) {
	data, err := ct.repo.GetStock(page, limit)
	if err != nil {
		return nil, err
	}

	totalStocks, err := ct.repo.CountStock()
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseStocks(data, totalStocks)
}

func (ct *StockUseCase) GetStockByID(ID int) (interface{}, error) {
	data, err := ct.repo.GetStockByID(ID)
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseStock(data)
}

func (ct *StockUseCase) UpdateStock(req service.UpdateStockRequest) error {
	var stock entity.Stock
	err := mapstructure.Decode(req, &stock)
	if err != nil {
		return nil
	}

	err = ct.repo.UpdateStock(stock)
	if err != nil {
		return err
	}

	return nil
}

func (ct *StockUseCase) DeleteStock(req service.DeleteStockRequest) error {
	var stock entity.Stock
	err := mapstructure.Decode(req, &stock)
	if err != nil {
		return err
	}

	err = ct.repo.DeleteStock(stock)
	if err != nil {
		return nil
	}

	return nil
}
