package main

import (
	"fmt"
	"group-project-2/configs"
	_userHandler "group-project-2/delivery/handler/user"
	_middleware "group-project-2/delivery/middlewares"
	_routes "group-project-2/delivery/routes"
	_userRepository "group-project-2/repository/user"
	_userUseCase "group-project-2/usecase/user"
	_utils "group-project-2/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middleware.CustomLogger())

	_routes.RegisterPathUser(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}
