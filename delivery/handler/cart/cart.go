package cart

import (
	"fmt"
	"group-project-2/delivery/helper"
	_middlewares "group-project-2/delivery/middlewares"
	_entities "group-project-2/entities"
	_cartUseCase "group-project-2/usecase/cart"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUseCase _cartUseCase.CartUseCaseInterface
}

func NewCartHandler(cartUseCase _cartUseCase.CartUseCaseInterface) *CartHandler {
	return &CartHandler{
		cartUseCase: cartUseCase,
	}
}
func (uh *CartHandler) PostCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, _ := _middlewares.ExtractToken(c)

		var cart _entities.Cart
		c.Bind(&cart)

		fmt.Print("ini di handler ", cart)
		cartNew, _, err := uh.cartUseCase.PostCart(cart, idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create Cart"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to insert new data", cartNew))

	}
}
