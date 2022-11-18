package presenter

import (
	"encoding/json"
	"pos/pkg/domain/entity"
	"pos/pkg/domain/service"
	"pos/pkg/shared/util"

	"github.com/mitchellh/mapstructure"
)

type SaleBuilder struct{}

func (*SaleBuilder) GetResponseSales(s interface{}, totalData int) (interface{}, error) {
	var sales []entity.Sale
	err := mapstructure.Decode(s, &sales)
	if err != nil {
		panic(err)
	}

	var Sales []service.Sale
	for i := range sales {
		var Sale service.Sale
		Sale.ID = sales[i].ID
		Sale.TotalPrice = sales[i].TotalPrice
		Sale.Discount = sales[i].Discount
		Sale.FinalPrice = sales[i].FinalPrice
		Sale.Cash = sales[i].Cash
		Sale.Remaining = sales[i].Remaining
		Sale.Note = sales[i].Note
		Sale.UserID = sales[i].UserID
		Sale.UpdatedAt = sales[i].UpdatedAt
		Sale.CreatedAt = sales[i].CreatedAt

		var jsonMapItem []util.PropertyMap
		err = json.Unmarshal([]byte(sales[i].Items), &jsonMapItem)
		Sale.Items = jsonMapItem

		Sales = append(Sales, Sale)
	}

	return service.SaleResponses{
		Sales: Sales,
		Total: totalData,
	}, nil
}

func (*SaleBuilder) GetResponseSale(u interface{}) (interface{}, error) {
	var sale entity.Sale
	err := mapstructure.Decode(u, &sale)
	if err != nil {
		return nil, err
	}

	var jsonMapItem []util.PropertyMap
	err = json.Unmarshal([]byte(sale.Items), &jsonMapItem)

	return service.SaleResponse{
		Sale: service.Sale{
			ID:         sale.ID,
			TotalPrice: sale.TotalPrice,
			Discount:   sale.Discount,
			FinalPrice: sale.FinalPrice,
			Cash:       sale.Cash,
			Remaining:  sale.Remaining,
			Note:       sale.Note,
			UserID:     sale.UserID,
			CreatedAt:  sale.CreatedAt,
			UpdatedAt:  sale.UpdatedAt,
			Items:      jsonMapItem,
		},
	}, nil
}
