package product

import (
	"fmt"
	"group-project-2/delivery/helper"
	_middleware "group-project-2/delivery/middlewares"
	_entities "group-project-2/entities"
	_productUseCase "group-project-2/usecase/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUseCase _productUseCase.ProductUseCaseInterface
}

func NewProductHandler(productUseCase _productUseCase.ProductUseCaseInterface) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

func (ph *ProductHandler) AddProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var product _entities.Product
		error := c.Bind(&product)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind data"))
		}
		product, err := ph.productUseCase.AddProduct(product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to add product"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to add new product", product))
	}
}

func (ph *ProductHandler) UpdateProductByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var product _entities.Product
		idToken, errToken := _middleware.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		id, _ := strconv.Atoi(c.Param("id"))
		err := c.Bind(&product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind data"))
		}
		_, rows, error := ph.productUseCase.UpdateProductById(product, id, idToken)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("product not found"))
		}
		successMessage := fmt.Sprintf("success to update product data by id = %v", id)
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData(successMessage))
	}
}

func (ph *ProductHandler) DeleteProductByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middleware.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		id, _ := strconv.Atoi(c.Param("id"))

		rows, err := ph.productUseCase.DeleteProductById(id, idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to operate delete"))
		}
		if rows == 0 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("product not found"))
		}
		successMessage := fmt.Sprintf("success to delete product data by id = %v", id)
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData(successMessage))
	}
}

func (ph *ProductHandler) GetAllProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := ph.productUseCase.GetAllProduct()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all books", products))
	}
}

func (ph *ProductHandler) GetProductByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("unrecognized id"))
		}
		product, rows, error := ph.productUseCase.GetProductById(id)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("product not found"))
		}
		if rows == 0 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("product not found"))
		}
		successMessage := fmt.Sprintf("success to get data from book id = %v", id)
		return c.JSON(http.StatusOK, helper.ResponseSuccess(successMessage, product))
	}
}
