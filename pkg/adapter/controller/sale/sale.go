package sale

import (
	"pos/pkg/domain/service"
	"pos/pkg/usecase/sale"

	"github.com/sirupsen/logrus"
)

type SaleService struct {
	SaleRepository sale.SaleInputPort
}

func NewSaleService(ct sale.SaleInputPort) *SaleService {
	return &SaleService{
		SaleRepository: ct,
	}
}
func (ct *SaleService) CreateSale(req service.CreateSaleRequest) error {
	err := ct.SaleRepository.CreateSale(req)
	if err != nil {
		return err
	}

	return nil
}

func (ct *SaleService) GetSale(page, limit int) (interface{}, error) {
	sale, err := ct.SaleRepository.GetSale(page, limit)
	if err != nil {
		logrus.Error("[GetSale] error get Sale in Controller")
		return nil, err
	}
	return sale, err
}

func (ct *SaleService) GetSaleByID(ID int) (interface{}, error) {
	sale, err := ct.SaleRepository.GetSaleByID(ID)
	if err != nil {
		logrus.Error("[GetSaleByID] error get sale by id in Controller")
		return nil, err
	}
	return sale, nil
}

func (ct *SaleService) DeleteSale(req service.DeleteSaleRequest) error {
	err := ct.SaleRepository.DeleteSale(req)
	if err != nil {
		return err
	}
	return err
}

func (ct *SaleService) UpdateSale(req service.UpdateSaleRequest) error {
	err := ct.SaleRepository.UpdateSale(req)
	if err != nil {
		return err
	}

	return nil
}
