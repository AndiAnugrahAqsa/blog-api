package routes

import (
	"mini-project/controllers"
	"mini-project/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var blogController = controllers.NewBlogController()
var categoryController = controllers.NewCategoryController()
var commentController = controllers.NewCommentController()
var likeController = controllers.NewLikeController()
var roleController = controllers.NewRoleController()
var userController = controllers.NewUserController()

func RoutesInit(e *echo.Echo) {

	superUserPrivateRoutes := e.Group("")

	configSuperUser := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKeyForSuperUser,
	}

	superUserPrivateRoutes.Use(middleware.JWTWithConfig(configSuperUser))

	userPrivateRoutes := e.Group("")

	configUser := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKeyForUser,
	}

	userPrivateRoutes.Use(middleware.JWTWithConfig(configUser))

	e.GET("/blogs", blogController.GetAll)
	e.GET("/blogs/user/:user_id", blogController.GetByUserID)
	e.GET("/blogs/category/:category_id", blogController.GetByCategoryID)
	e.GET("/blogs/:id", blogController.GetByID)
	userPrivateRoutes.POST("/blogs", blogController.Create)
	userPrivateRoutes.PUT("/blogs/:id", blogController.Update)
	userPrivateRoutes.DELETE("/blogs/:id", blogController.Delete)

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

	e.GET("/likes", likeController.GetAll)
	e.GET("likes/blog/:blog_id", likeController.GetByBlogID)
	userPrivateRoutes.POST("/likes", likeController.Create)
	userPrivateRoutes.DELETE("/likes/:id", likeController.Delete)

	superUserPrivateRoutes.GET("/roles", roleController.GetAll)
	superUserPrivateRoutes.GET("/roles/:id", roleController.GetByID)
	superUserPrivateRoutes.POST("/roles", roleController.Create)
	superUserPrivateRoutes.PUT("/roles/:id", roleController.Update)
	superUserPrivateRoutes.DELETE("/roles/:id", roleController.Delete)

	superUserPrivateRoutes.GET("/users", userController.GetAll)
	superUserPrivateRoutes.GET("/users/:id", userController.GetByID)
	superUserPrivateRoutes.POST("/users", userController.Create)
	superUserPrivateRoutes.PUT("/users/:id", userController.Update)
	superUserPrivateRoutes.DELETE("/users/:id", userController.Delete)

	// user auth
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)
	e.POST("/logout", userController.Logout)
}
