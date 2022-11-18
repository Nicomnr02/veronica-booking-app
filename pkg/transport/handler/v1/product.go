package v1

import (
	"net/http"
	pr "pos/pkg/adapter/controller/product"
	"pos/pkg/domain/service"
	"pos/pkg/shared/util"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

func CreateProduct(svc pr.ProductService) func(echo.Context) error {
	return func(c echo.Context) error {
		var Product service.CreateProductRequest
		c.Bind(&Product)

		err := svc.CreateProduct(Product)

		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[CreateProduct] error when create Product", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func UpdateProduct(svc pr.ProductService) func(echo.Context) error {
	return func(c echo.Context) error {
		var ProductPayload service.UpdateProductRequest

		err := c.Bind(&ProductPayload)
		if err != nil {
			return util.SetResponse(c, http.StatusBadRequest, "[UpdateProduct] error : check input data type", nil)
		}

		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "ic can't be blank", nil)
		}

		ProductPayload.ID = id

		// err2 := service.UpdateProductRequest.Validate(ProductPayload)
		// if err2 != nil {
		// 	return util.SetResponse(c, http.StatusBadRequest, err2.Error(), nil)
		// }

		err3 := svc.UpdateProduct(ProductPayload)
		if err3 != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[UpdateProduct] error when update Product", nil)
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func GetProduct(svc pr.ProductService) func(echo.Context) error {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		ss, err := svc.GetProduct(page, limit)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetProduct] error when Get Product ", nil)
		}
		return util.SetResponse(c, http.StatusOK, "success", ss)
	}
}

func GetProductByID(svc pr.ProductService) func(echo.Context) error {
	return func(c echo.Context) error {
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		ss, err := svc.GetProductByID(Id)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[GetProductByID] error when get croduct by id", nil)
		}

		return util.SetResponse(c, http.StatusOK, "succes", ss)
	}
}

func DeleteProduct(svc pr.ProductService) func(echo.Context) error {
	return func(c echo.Context) error {
		var ProductPayload service.DeleteProductRequest
		Id, _ := strconv.Atoi(c.Param("id"))
		if Id == 0 {
			return util.SetResponse(c, http.StatusBadRequest, "id can't be blank", nil)
		}
		ProductPayload.ID = Id

		err := svc.DeleteProduct(ProductPayload)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, "[DeleteProduct] error when delete croduct", nil)
		}
		return util.SetResponse(c, http.StatusOK, "Success", nil)
	}
}
