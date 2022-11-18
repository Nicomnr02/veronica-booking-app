package group

import (
	ct "pos/pkg/adapter/controller/product"
	v1 "pos/pkg/transport/handler/v1"

	echo "github.com/labstack/echo/v4"
)

func InitProductV1(e *echo.Echo, g *echo.Group, svc *ct.ProductService) {
	e.POST("/v1/product", v1.CreateProduct(*svc))
	e.GET("/v1/product", v1.GetProduct(*svc))
	e.GET("/v1/product/:id", v1.GetProductByID(*svc))
	e.PUT("/v1/product/:id", v1.UpdateProduct(*svc))
	e.DELETE("/v1/product/:id", v1.DeleteProduct(*svc))
}
