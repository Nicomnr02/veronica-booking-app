package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	cfg "pos/internal"
	container "pos/pkg/shared/di"
	"pos/pkg/transport/group"
	"pos/pkg/usecase/category"
	"pos/pkg/usecase/product"
	"pos/pkg/usecase/sale"
	"pos/pkg/usecase/size"
	"pos/pkg/usecase/stock"

	ct "pos/pkg/adapter/controller/category"
	pr "pos/pkg/adapter/controller/product"
	sl "pos/pkg/adapter/controller/sale"
	sz "pos/pkg/adapter/controller/size"
	st "pos/pkg/adapter/controller/stock"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RunServer() {
	configuration := cfg.GetConfig()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://localhost:8081",
			"https://pos.shineshop.com",
		},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// using JWT
	jwtGroup := e.Group("")
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("h1C01134guE5"),
		TokenLookup:   "header:Authorization",
		AuthScheme:    "Bearer",
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	ctn := container.NewContainer(configuration)
	Apply(ctn, e, jwtGroup)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", configuration.Server.HC.Port)))
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal %v\n", signal.String())
}

func Apply(ctn *container.Container, e *echo.Echo, g *echo.Group) {
	categorySvc := ct.NewCategoryService(ctn.Resolve("categorysvc").(*category.CategoryUseCase))
	sizeSvc := sz.NewSizeService(ctn.Resolve("sizesvc").(*size.SizeUseCase))
	productSvc := pr.NewProductService(ctn.Resolve("productsvc").(*product.ProductUseCase))
	stockSvc := st.NewStockService(ctn.Resolve("stocksvc").(*stock.StockUseCase))
	saleSvc := sl.NewSaleService(ctn.Resolve("salesvc").(*sale.SaleUseCase))

	group.InitCategoryV1(e, g, categorySvc)
	group.InitSizeV1(e, g, sizeSvc)
	group.InitProductV1(e, g, productSvc)
	group.InitStockV1(e, g, stockSvc)
	group.InitSaleV1(e, g, saleSvc)

}
