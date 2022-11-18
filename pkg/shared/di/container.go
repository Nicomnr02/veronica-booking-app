package di

import (
	cfg "pos/internal"
	"pos/pkg/adapter/presenter"

	// "pos/pkg/adapter/repository/api/masterdata"
	db "pos/pkg/adapter/repository/document"
	"pos/pkg/usecase/category"
	"pos/pkg/usecase/product"
	"pos/pkg/usecase/sale"
	"pos/pkg/usecase/size"
	"pos/pkg/usecase/stock"

	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

var config *cfg.Config

func NewContainer(configuration *cfg.Config) *Container {
	builder, _ := di.NewBuilder()
	config = configuration

	_ = builder.Add([]di.Def{
		{Name: "categorysvc", Build: categoryUsecase},
		{Name: "sizesvc", Build: sizeUsecase},
		{Name: "productsvc", Build: productUsecase},
		{Name: "stocksvc", Build: stockUsecase},
		{Name: "salesvc", Build: saleUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func categoryUsecase(_ di.Container) (interface{}, error) {
	categoryBuilder := &presenter.CategoryBuilder{}
	categoryRepo := db.NewCategory(config.Database)

	return category.NewCategoryUseCase(categoryRepo, categoryBuilder), nil
}

func sizeUsecase(_ di.Container) (interface{}, error) {
	sizeBuilder := &presenter.SizeBuilder{}
	sizeRepo := db.NewSize(config.Database)

	return size.NewSizeUseCase(sizeRepo, sizeBuilder), nil
}

func productUsecase(_ di.Container) (interface{}, error) {
	productBuilder := &presenter.ProductBuilder{}
	productRepo := db.NewProduct(config.Database)

	return product.NewProductUseCase(productRepo, productBuilder), nil
}

func stockUsecase(_ di.Container) (interface{}, error) {
	stockBuilder := &presenter.StockBuilder{}
	stockRepo := db.NewStock(config.Database)

	return stock.NewStockUseCase(stockRepo, stockBuilder), nil
}

func saleUsecase(_ di.Container) (interface{}, error) {
	saleBuilder := &presenter.SaleBuilder{}
	saleRepo := db.NewSale(config.Database)

	return sale.NewSaleUseCase(saleRepo, saleBuilder), nil
}
