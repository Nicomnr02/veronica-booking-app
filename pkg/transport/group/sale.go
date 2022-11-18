package group

import (
	ct "pos/pkg/adapter/controller/sale"
	v1 "pos/pkg/transport/handler/v1"

	echo "github.com/labstack/echo/v4"
)

func InitSaleV1(e *echo.Echo, g *echo.Group, svc *ct.SaleService) {
	e.POST("/v1/sale", v1.CreateSale(*svc))
	e.GET("/v1/sale", v1.GetSale(*svc))
	e.GET("/v1/sale/:id", v1.GetSaleByID(*svc))
	e.PUT("/v1/sale/:id", v1.UpdateSale(*svc))
	e.DELETE("/v1/sale/:id", v1.DeleteSale(*svc))
}
