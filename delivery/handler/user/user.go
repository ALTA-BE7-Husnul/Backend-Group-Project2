package user

import (
	"group-project-2/delivery/helper"
	_middlewares "group-project-2/delivery/middlewares"
	_entities "group-project-2/entities"
	_userUseCase "group-project-2/usecase/user"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(userUseCase _userUseCase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}
func (uh *UserHandler) PostUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.User
		tx := c.Bind(&user)
		users, _ := uh.userUseCase.GetAll()
		for _, val := range users {
			if val.Email == user.Email {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Email already exist"))
			}
		}

		userNew, _ := uh.userUseCase.PostUser(user)
		if tx != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create User"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to insert new data", userNew))

	}
}
func (uh *UserHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch user"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses get all data", users))
	}
}
func (uh *UserHandler) GetUserHandler() echo.HandlerFunc {
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
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized or different users"))
		}

		users, rows, err := uh.userUseCase.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("User does not exist"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to get data by ID", users))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
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
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized or different users"))
		}

		users, rows, err := uh.userUseCase.DeleteUser(id)
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete User"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to delete User", users))
	}
}

func (uh *UserHandler) PutUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		var user _entities.User
		var updateUser _entities.User

		tx := c.Bind(&updateUser)
		if tx != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update User: Check format request body"))
		}
		if updateUser.Name != "" {
			user.Name = updateUser.Name
		}
		if updateUser.Email != "" {
			user.Email = updateUser.Email
		}
		if updateUser.Address != "" {
			user.Address = updateUser.Address
		}
		if updateUser.Password != "" {
			user.Password = updateUser.Password
		}

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		if idToken != int(user.ID) {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized or different users"))
		}

		user, _, err := uh.userUseCase.PutUser(idToken, user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to update user"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to update User", user))
	}
}
