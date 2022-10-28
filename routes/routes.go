package routes

import (
	"mini-project/controllers"
	"mini-project/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var roleController = controllers.NewRoleController()
var userController = controllers.NewUserController()

func RoutesInit(e *echo.Echo) {

	superUserPrivateRoutes := e.Group("")

	config := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKey,
	}

	superUserPrivateRoutes.Use(middleware.JWTWithConfig(config))

	superUserPrivateRoutes.GET("/roles", roleController.GetAll)
	superUserPrivateRoutes.GET("/roles/:id", roleController.GetByID)
	superUserPrivateRoutes.POST("/roles", roleController.Create)
	superUserPrivateRoutes.PUT("/roles/:id", roleController.Update)
	superUserPrivateRoutes.DELETE("/roles/:id", roleController.Delete)

	superUserPrivateRoutes.GET("/users", userController.GetAll)
	superUserPrivateRoutes.GET("/users/:id", userController.GetByID)
	superUserPrivateRoutes.PUT("/users/:id", userController.Update)
	superUserPrivateRoutes.DELETE("/users/:id", userController.Delete)

	// user auth
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)
}
