package order

import (
	"group-project-2/delivery/helper"
	_helper "group-project-2/delivery/helper"
	_middleware "group-project-2/delivery/middlewares"
	_orderUseCase "group-project-2/usecase/order"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUseCase _orderUseCase.OrderUseCaseInterface
}

func NewOrderHandler(orderUseCase _orderUseCase.OrderUseCaseInterface) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

func (oh *OrderHandler) GetOrderHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middleware.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		orders, err := oh.orderUseCase.GetOrder(idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all order", orders))
	}
}

func (oh *OrderHandler) PostOrderHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var order _helper.OrderRequestFormat
		_, errToken := _middleware.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		bindError := c.Bind(&order)
		if bindError != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind address data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success make an order", order))
	}
}
