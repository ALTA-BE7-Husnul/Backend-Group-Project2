package cart

import (
	"group-project-2/delivery/helper"
	_middlewares "group-project-2/delivery/middlewares"
	_entities "group-project-2/entities"
	_cartUseCase "group-project-2/usecase/cart"
	"net/http"
	"strconv"

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
		cartNew, _, rows, err := uh.cartUseCase.PostCart(cart, idToken)

		if rows == 1 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("not enaough product"))
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create Cart"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to insert new data", cartNew))

	}
}
func (uh *CartHandler) GetCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		cart, err := uh.cartUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch cart"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses get all data", cart))
	}
}
func (ph *CartHandler) PutCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		var cart _entities.Cart
		var updateCart _entities.Cart

		c.Bind(&updateCart)

		idStr := c.Param("id")
		idToken, errorconv := strconv.Atoi(idStr)
		if errorconv != nil {
			return c.JSON(http.StatusBadRequest, "The expected param must be int")
		}

		if updateCart.Buyer_ID == 0 {
			cart.Buyer_ID = updateCart.Buyer_ID
		}
		// if updateCart.Status != "" {
		// 	cart.Status = updateCart.Status
		// }

		cart, err := ph.cartUseCase.PutCart(updateCart, idToken)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to update cart"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to update Cart", cart))
	}
}
func (uh *CartHandler) DeleteCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idStr := c.Param("id")
		id, errorconv := strconv.Atoi(idStr)
		if errorconv != nil {
			return c.JSON(http.StatusBadRequest, "The expected param must be int")
		}

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized or different carts"))
		}

		carts, rows, err := uh.cartUseCase.DeleteCart(id)
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete Cart"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to delete Cart", carts))
	}
}
