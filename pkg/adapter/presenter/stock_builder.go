package presenter

import (
	"pos/pkg/domain/entity"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type StockBuilder struct{}

func (*StockBuilder) GetResponseStocks(s interface{}, totalData int) (interface{}, error) {

	var stocks []entity.Stock
	err := mapstructure.Decode(s, &stocks)
	if err != nil {
		panic(err)
	}

	return service.StockResponses{
		Stocks: stocks,
		Total:  totalData,
	}, nil
}

func (*StockBuilder) GetResponseStock(u interface{}) (interface{}, error) {
	var stock entity.Stock
	err := mapstructure.Decode(u, &stock)
	if err != nil {
		return nil, err
	}

	return service.StockResponse{
		Stock: stock,
	}, nil
}
