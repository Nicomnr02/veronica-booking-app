package sale

import (
	"encoding/json"
	"fmt"
	"pos/pkg/domain/entity"
	"pos/pkg/domain/repository"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type SaleUseCase struct {
	repo repository.SaleRepository
	out  SaleOutputPort
}

func NewSaleUseCase(r repository.SaleRepository, o SaleOutputPort) *SaleUseCase {
	return &SaleUseCase{
		repo: r,
		out:  o,
	}
}

func (ct *SaleUseCase) CreateSale(req service.CreateSaleRequest) error {
	var sale entity.Sale

	itemsMap, _ := json.Marshal(req.Items)
	sale = entity.Sale{
		TotalPrice: req.TotalPrice,
		Discount:   req.Discount,
		FinalPrice: req.FinalPrice,
		Cash:       req.Cash,
		Remaining:  req.Remaining,
		Note:       req.Note,
		UserID:     req.UserID,
		Items:      string(itemsMap),
	}

	err := ct.repo.CreateSale(sale)
	if err != nil {
		return err
	}
	return nil
}

func (ct *SaleUseCase) GetSale(page, limit int) (interface{}, error) {
	data, err := ct.repo.GetSale(page, limit)
	if err != nil {
		return nil, err
	}

	totalCategories, err := ct.repo.CountSale()
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseSales(data, totalCategories)
}

func (ct *SaleUseCase) GetSaleByID(ID int) (interface{}, error) {
	data, err := ct.repo.GetSaleByID(ID)
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseSale(data)
}

func (ct *SaleUseCase) UpdateSale(req service.UpdateSaleRequest) error {
	var sale entity.Sale

	itemsMap, _ := json.Marshal(req.Items)
	sale = entity.Sale{
		ID:         req.ID,
		TotalPrice: req.TotalPrice,
		Discount:   req.Discount,
		FinalPrice: req.FinalPrice,
		Cash:       req.Cash,
		Remaining:  req.Remaining,
		Note:       req.Note,
		UserID:     req.UserID,
		Items:      string(itemsMap),
	}

	err := ct.repo.UpdateSale(sale)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

func (ct *SaleUseCase) DeleteSale(req service.DeleteSaleRequest) error {
	var sale entity.Sale
	err := mapstructure.Decode(req, &sale)
	if err != nil {
		return err
	}

	err = ct.repo.DeleteSale(sale)
	if err != nil {
		return err
	}

	return nil
}
