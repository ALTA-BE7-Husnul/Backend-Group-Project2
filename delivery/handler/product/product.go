package product

import (
	"group-project-2/delivery/helper"
	"group-project-2/entities"
	_productUseCase "group-project-2/usecase/product"
	"net/http"

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
		var product entities.Product
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
