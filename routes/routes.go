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

	adminPrivateRoutes := e.Group("")

	configAdmin := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKeyForAdmin,
	}

	adminPrivateRoutes.Use(middleware.JWTWithConfig(configAdmin))

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
	adminPrivateRoutes.POST("/categories", categoryController.Create)
	adminPrivateRoutes.PUT("/categories/:id", categoryController.Update)
	adminPrivateRoutes.DELETE("/categories/:id", categoryController.Delete)

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

	adminPrivateRoutes.GET("/roles", roleController.GetAll)
	adminPrivateRoutes.GET("/roles/:id", roleController.GetByID)
	adminPrivateRoutes.POST("/roles", roleController.Create)
	adminPrivateRoutes.PUT("/roles/:id", roleController.Update)
	adminPrivateRoutes.DELETE("/roles/:id", roleController.Delete)

	adminPrivateRoutes.GET("/users", userController.GetAll)
	adminPrivateRoutes.GET("/users/:id", userController.GetByID)
	adminPrivateRoutes.POST("/users", userController.Create)
	adminPrivateRoutes.PUT("/users/:id", userController.Update)
	adminPrivateRoutes.DELETE("/users/:id", userController.Delete)

	// user auth
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)
	e.POST("/logout", userController.Logout)
}
