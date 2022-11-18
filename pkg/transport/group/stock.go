package group

import (
	ct "pos/pkg/adapter/controller/stock"
	v1 "pos/pkg/transport/handler/v1"

	echo "github.com/labstack/echo/v4"
)

func InitStockV1(e *echo.Echo, g *echo.Group, svc *ct.StockService) {
	e.POST("/v1/stock", v1.CreateStock(*svc))
	e.GET("/v1/stock", v1.GetStock(*svc))
	e.GET("/v1/stock/:id", v1.GetStockByID(*svc))
	e.PUT("/v1/stock/:id", v1.UpdateStock(*svc))
	e.DELETE("/v1/stock/:id", v1.DeleteStock(*svc))
}
