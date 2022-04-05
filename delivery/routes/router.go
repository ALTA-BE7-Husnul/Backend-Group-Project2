package routes

import (
	_authHandler "group-project-2/delivery/handler/auth"
	_userHandler "group-project-2/delivery/handler/user"
	_middlewares "group-project-2/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login", ah.LoginHandler())
}
func RegisterPathUser(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users/", uh.PostUserHandler())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.PutUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())

}
