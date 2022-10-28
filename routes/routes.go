package routes

import (
	"mini-project/controllers"

	"github.com/labstack/echo/v4"
)

var roleController = controllers.NewRoleController()
var userController = controllers.NewUserController()

func RoutesInit(e *echo.Echo) {
	e.GET("/roles", roleController.GetAll)
	e.GET("/roles/:id", roleController.GetByID)
	e.POST("/roles", roleController.Create)
	e.PUT("/roles/:id", roleController.Update)
	e.DELETE("/roles/:id", roleController.Delete)

	e.GET("/users", userController.GetAll)
	e.GET("/users/:id", userController.GetByID)
	e.PUT("/users/:id", userController.Update)
	e.DELETE("/users/:id", userController.Delete)

	// user auth
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)
}
