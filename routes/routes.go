package routes

import (
	"mini-project/controllers"
	"mini-project/middlewares"
	"mini-project/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var categoryController = controllers.NewCategoryController()
var commentController = controllers.NewCommentController()
var roleController = controllers.NewRoleController()
var userController = controllers.NewUserController()

func RoutesInit(e *echo.Echo) {

	superUserPrivateRoutes := e.Group("")

	config := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKey,
	}

	superUserPrivateRoutes.Use(middleware.JWTWithConfig(config))

	userPrivateRoutes := e.Group("")

	userPrivateRoutes.Use(middleware.JWT([]byte(util.GetConfig("JWT_SECRET_KEY"))))

	e.GET("/categories", categoryController.GetAll)
	e.GET("/categories/:id", categoryController.GetByID)
	superUserPrivateRoutes.POST("/categories", categoryController.Create)
	superUserPrivateRoutes.PUT("/categories/:id", categoryController.Update)
	superUserPrivateRoutes.DELETE("/categories/:id", categoryController.Delete)

	e.GET("/comments", commentController.GetAll)
	e.GET("/comments/:id", commentController.GetByID)
	e.GET("comments/blog/:blog_id", commentController.GetByBlogID)
	userPrivateRoutes.POST("/comments", commentController.Create)
	userPrivateRoutes.PUT("/comments/:id", commentController.Update)
	userPrivateRoutes.DELETE("/comments/:id", commentController.Delete)

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
