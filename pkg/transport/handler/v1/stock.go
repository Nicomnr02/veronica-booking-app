package v1

import (
	"fmt"
	"net/http"
	st "pos/pkg/adapter/controller/stock"
	"pos/pkg/domain/service"
	"pos/pkg/shared/util"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

func CreateStock(svc st.StockService) func(echo.Context) error {
	return func(c echo.Context) error {
		var Stock service.CreateStockRequest
		c.Bind(&Stock)
		fmt.Println(Stock)
		err := svc.CreateStock(Stock)

		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[CreateStock] error when create Stock", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func UpdateStock(svc st.StockService) func(echo.Context) error {
	return func(c echo.Context) error {
		var StockPayload service.UpdateStockRequest

		err := c.Bind(&StockPayload)
		if err != nil {
			return util.SetResponse(c, http.StatusBadRequest, "[UpdateStock] error : check input data type", nil)
		}

		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "ic can't be blank", nil)
		}

		StockPayload.ID = id

		// err2 := service.UpdateStockRequest.Validate(StockPayload)
		// if err2 != nil {
		// 	return util.SetResponse(c, http.StatusBadRequest, err2.Error(), nil)
		// }

		err3 := svc.UpdateStock(StockPayload)
		if err3 != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[UpdateStock] error when update Stock", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func GetStock(svc st.StockService) func(echo.Context) error {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		ss, err := svc.GetStock(page, limit)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetStock] error when Get Stock ", nil)
		}
		return util.SetResponse(c, http.StatusOK, "success", ss)
	}
}

func GetStockByID(svc st.StockService) func(echo.Context) error {
	return func(c echo.Context) error {
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		ss, err := svc.GetStockByID(Id)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetStockByID] error when get croduct by id", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", ss)
	}
}

func DeleteStock(svc st.StockService) func(echo.Context) error {
	return func(c echo.Context) error {
		var StockPayload service.DeleteStockRequest
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		StockPayload.ID = Id

		err := svc.DeleteStock(StockPayload)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[DeleteStock] error when delete croduct", nil)
		}
		return util.SetResponse(c, http.StatusOK, "Success", nil)
	}
}
