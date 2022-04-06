package main

import (
	"fmt"
	"group-project-2/configs"
	_authHandler "group-project-2/delivery/handler/auth"
	_productHandler "group-project-2/delivery/handler/product"
	_userHandler "group-project-2/delivery/handler/user"
	_middleware "group-project-2/delivery/middlewares"
	_routes "group-project-2/delivery/routes"
	_authRepository "group-project-2/repository/auth"
	_productRepository "group-project-2/repository/product"
	_userRepository "group-project-2/repository/user"
	_authUseCase "group-project-2/usecase/auth"
	_productUseCase "group-project-2/usecase/product"
	_userUseCase "group-project-2/usecase/user"
	_utils "group-project-2/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	productRepo := _productRepository.NewProductRepository(db)
	productUseCase := _productUseCase.NewProductUseCase(productRepo)
	productHandler := _productHandler.NewProductHandler(productUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middleware.CustomLogger())

	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterPathUser(e, userHandler)
	_routes.RegisterPathProduct(e, productHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}
