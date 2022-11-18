package group

import (
	sz "pos/pkg/adapter/controller/size"
	v1 "pos/pkg/transport/handler/v1"

	echo "github.com/labstack/echo/v4"
)

func InitSizeV1(e *echo.Echo, g *echo.Group, svc *sz.SizeService) {
	e.POST("/v1/size", v1.CreateSize(*svc))
	e.GET("/v1/size", v1.GetSize(*svc))
	e.GET("/v1/size/:id", v1.GetSizeByID(*svc))
	e.PUT("/v1/size/:id", v1.UpdateSize(*svc))
	e.DELETE("/v1/size/:id", v1.DeleteSize(*svc))
}
