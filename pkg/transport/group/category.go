package group

import (
	ct "pos/pkg/adapter/controller/category"
	v1 "pos/pkg/transport/handler/v1"

	echo "github.com/labstack/echo/v4"
)

func InitCategoryV1(e *echo.Echo, g *echo.Group, svc *ct.CategoryService) {
	e.POST("/v1/category", v1.CreateCategory(*svc))
	e.GET("/v1/category", v1.GetCategory(*svc))
	e.GET("/v1/category/:id", v1.GetCategoryByID(*svc))
	e.PUT("/v1/category/:id", v1.UpdateCategory(*svc))
	e.DELETE("/v1/category/:id", v1.DeleteCategory(*svc))
}
