package routes

import (
	"mini-project/controllers"

	"github.com/labstack/echo/v4"
)

var roleController = controllers.NewRoleController()

func RoutesInit(e *echo.Echo) {
	e.GET("/roles", roleController.GetAll)
	e.GET("/roles/:id", roleController.GetByID)
	e.POST("/roles/:id", roleController.Create)
	e.PUT("/roles/:id", roleController.Update)
	e.DELETE("/roles/:id", roleController.Delete)
}
