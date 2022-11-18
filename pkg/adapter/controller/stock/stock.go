package stock

import (
	"pos/pkg/domain/service"
	"pos/pkg/usecase/stock"

	"github.com/sirupsen/logrus"
)

type StockService struct {
	StockRepository stock.StockInputPort
}

func NewStockService(ct stock.StockInputPort) *StockService {
	return &StockService{
		StockRepository: ct,
	}
}
func (ct *StockService) CreateStock(req service.CreateStockRequest) error {
	err := ct.StockRepository.CreateStock(req)
	if err != nil {
		return err
	}

	return nil
}

func (ct *StockService) GetStock(page, limit int) (interface{}, error) {
	stock, err := ct.StockRepository.GetStock(page, limit)
	if err != nil {
		logrus.Error("[GetStock] error get Stock in Controller")
		return nil, err
	}
	return stock, err
}

func (ct *StockService) GetStockByID(ID int) (interface{}, error) {
	stock, err := ct.StockRepository.GetStockByID(ID)
	if err != nil {
		logrus.Error("[GetStockByID] error get stock by id in Controller")
		return nil, err
	}
	return stock, nil
}

func (ct *StockService) DeleteStock(req service.DeleteStockRequest) error {
	err := ct.StockRepository.DeleteStock(req)
	if err != nil {
		return nil
	}
	return err
}

func (ct *StockService) UpdateStock(req service.UpdateStockRequest) error {
	err := ct.StockRepository.UpdateStock(req)
	if err != nil {
		return err
	}

	return nil
}
