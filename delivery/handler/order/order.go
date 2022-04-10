package order

import (
	"group-project-2/delivery/helper"
	_middleware "group-project-2/delivery/middlewares"
	"group-project-2/entities"
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
		var orderRequest helper.OrderRequestFormat
		idToken, errToken := _middleware.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		bindErr := c.Bind(&orderRequest)
		if bindErr != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind order request"))
		}

		addressRequest := entities.Address{
			Street:  orderRequest.Address.Street,
			City:    orderRequest.Address.City,
			State:   orderRequest.Address.State,
			Zipcode: orderRequest.Address.Zipcode,
		}

		paymentRequest := entities.Payment{
			Method:          orderRequest.Payment.Method,
			DestinationBank: orderRequest.Payment.DestinationBank,
		}
		order := entities.Transaction{
			User_ID: uint(idToken),
			Address: addressRequest,
			Payment: paymentRequest,
		}

		orderCartID := orderRequest.Cart_ID

		_, rows, err := oh.orderUseCase.PostOrder(order, orderCartID, idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success make an order"))
	}
}
