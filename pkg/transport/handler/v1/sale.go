package v1

import (
	"net/http"
	ct "pos/pkg/adapter/controller/sale"
	"pos/pkg/domain/service"
	"pos/pkg/shared/util"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

func CreateSale(svc ct.SaleService) func(echo.Context) error {
	return func(c echo.Context) error {
		var Sale service.CreateSaleRequest
		c.Bind(&Sale)

		err := svc.CreateSale(Sale)

		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[CreateSale] error when create Sale", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func UpdateSale(svc ct.SaleService) func(echo.Context) error {
	return func(c echo.Context) error {
		var SalePayload service.UpdateSaleRequest

		err := c.Bind(&SalePayload)
		if err != nil {
			return util.SetResponse(c, http.StatusBadRequest, "[UpdateSale] error : check input data type", nil)
		}

		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "ic can't be blank", nil)
		}

		SalePayload.ID = id

		// entar aku buat yaa validasi nya

		// err2 := service.UpdateSaleRequest.Validate(SalePayload)
		// if err2 != nil {
		// 	return util.SetResponse(c, http.StatusBadRequest, err2.Error(), nil)
		// }

		err3 := svc.UpdateSale(SalePayload)
		if err3 != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[UpdateSale] error when update Sale", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func GetSale(svc ct.SaleService) func(echo.Context) error {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		ss, err := svc.GetSale(page, limit)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetSale] error when Get Sale ", nil)
		}
		return util.SetResponse(c, http.StatusOK, "success", ss)
	}
}

func GetSaleByID(svc ct.SaleService) func(echo.Context) error {
	return func(c echo.Context) error {
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		ss, err := svc.GetSaleByID(Id)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetSaleByID] error when get category by id", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", ss)
	}
}

func DeleteSale(svc ct.SaleService) func(echo.Context) error {
	return func(c echo.Context) error {
		var SalePayload service.DeleteSaleRequest
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		SalePayload.ID = Id

		err := svc.DeleteSale(SalePayload)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[DeleteSale] error when delete category", nil)
		}
		return util.SetResponse(c, http.StatusOK, "Success", nil)
	}
}
