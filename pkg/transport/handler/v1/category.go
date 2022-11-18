package v1

import (
	"net/http"
	ct "pos/pkg/adapter/controller/category"
	"pos/pkg/domain/service"
	"pos/pkg/shared/util"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

func CreateCategory(svc ct.CategoryService) func(echo.Context) error {
	return func(c echo.Context) error {
		var Category service.CreateCategoryRequest
		c.Bind(&Category)

		err := svc.CreateCategory(Category)

		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[CreateCategory] error when create Category", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func UpdateCategory(svc ct.CategoryService) func(echo.Context) error {
	return func(c echo.Context) error {
		var CategoryPayload service.UpdateCategoryRequest

		err := c.Bind(&CategoryPayload)
		if err != nil {
			return util.SetResponse(c, http.StatusBadRequest, "[UpdateCategory] error : check input data type", nil)
		}

		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "ic can't be blank", nil)
		}

		CategoryPayload.ID = id

		err2 := service.UpdateCategoryRequest.Validate(CategoryPayload)
		if err2 != nil {
			return util.SetResponse(c, http.StatusBadRequest, err2.Error(), nil)
		}

		err3 := svc.UpdateCategory(CategoryPayload)
		if err3 != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[UpdateCategory] error when update Category", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func GetCategory(svc ct.CategoryService) func(echo.Context) error {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		ss, err := svc.GetCategory(page, limit)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetCategory] error when Get Category ", nil)
		}
		return util.SetResponse(c, http.StatusOK, "success", ss)
	}
}

func GetCategoryByID(svc ct.CategoryService) func(echo.Context) error {
	return func(c echo.Context) error {
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		ss, err := svc.GetCategoryByID(Id)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetCategoryByID] error when get category by id", nil)
		}

		return util.SetResponse(c, http.StatusOK, "succes", ss)
	}
}

func DeleteCategory(svc ct.CategoryService) func(echo.Context) error {
	return func(c echo.Context) error {
		var CategoryPayload service.DeleteCategoryRequest
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		CategoryPayload.ID = Id

		err := svc.DeleteCategory(CategoryPayload)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[DeleteCategory] error when delete category", nil)
		}
		return util.SetResponse(c, http.StatusOK, "Success", nil)
	}
}
