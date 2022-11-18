package v1

import (
	"net/http"
	sz "pos/pkg/adapter/controller/size"
	"pos/pkg/domain/service"
	"pos/pkg/shared/util"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

func CreateSize(svc sz.SizeService) func(echo.Context) error {
	return func(c echo.Context) error {
		var Size service.CreateSizeRequest
		c.Bind(&Size)

		err := svc.CreateSize(Size)

		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[CreateSize] error when create Size", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func UpdateSize(svc sz.SizeService) func(echo.Context) error {
	return func(c echo.Context) error {
		var SizePayload service.UpdateSizeRequest

		err := c.Bind(&SizePayload)
		if err != nil {
			return util.SetResponse(c, http.StatusBadRequest, "[UpdateSize] error : check input data type", nil)
		}

		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "ic can't be blank", nil)
		}

		SizePayload.ID = id

		// err2 := service.UpdateSizeRequest.Validate(SizePayload)
		// if err2 != nil {
		// 	return util.SetResponse(c, http.StatusBadRequest, err2.Error(), nil)
		// }

		err3 := svc.UpdateSize(SizePayload)
		if err3 != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[UpdateSize] error when update Size", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func GetSize(svc sz.SizeService) func(echo.Context) error {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		ss, err := svc.GetSize(page, limit)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetSize] error when Get Size ", nil)
		}
		return util.SetResponse(c, http.StatusOK, "success", ss)
	}
}

func GetSizeByID(svc sz.SizeService) func(echo.Context) error {
	return func(c echo.Context) error {
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		ss, err := svc.GetSizeByID(Id)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetSizeByID] error when get Size by id", nil)
		}

		return util.SetResponse(c, http.StatusOK, "succes", ss)
	}
}

func DeleteSize(svc sz.SizeService) func(echo.Context) error {
	return func(c echo.Context) error {
		var SizePayload service.DeleteSizeRequest
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		SizePayload.ID = Id

		err := svc.DeleteSize(SizePayload)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[DeleteSize] error when delete Size", nil)
		}
		return util.SetResponse(c, http.StatusOK, "Success", nil)
	}
}
