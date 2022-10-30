package controllers

import (
	"mini-project/models"
	"mini-project/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	Service services.BlogService
}

func NewBlogController() BlogController {
	return BlogController{
		Service: services.NewBlogService(),
	}
}

func (bc *BlogController) GetAll(c echo.Context) error {
	var blogs []models.Blog
	blogs = bc.Service.Repository.GetAll()

	var blogsResponse []models.BlogResponse

	for _, blog := range blogs {
		blogsResponse = append(blogsResponse, blog.ToResponse())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully get all blogs",
		"blogs":   blogsResponse,
	})
}

func (bc *BlogController) GetByUserID(c echo.Context) error {
	userIDString := c.Param("user_id")
	userID, _ := strconv.Atoi(userIDString)

	var blogs []models.Blog
	blogs = bc.Service.Repository.GetByUserID(userID)

	var blogsResponse []models.BlogResponse

	for _, blog := range blogs {
		blogsResponse = append(blogsResponse, blog.ToResponse())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully get blogs by user id",
		"blogs":   blogsResponse,
	})
}

func (bc *BlogController) GetByID(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var blog models.Blog

	blog = bc.Service.Repository.GetByID(id)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully get blog",
		"blogs":   blog,
	})
}

func (bc *BlogController) Create(c echo.Context) error {
	var blogRequest models.BlogRequest

	c.Bind(&blogRequest)

	blog := bc.Service.Repository.Create(blogRequest)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully create blog",
		"blog":    blog.ToResponse(),
	})
}

func (bc *BlogController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var blogUpdate models.BlogRequest

	c.Bind(&blogUpdate)

	blog := bc.Service.Repository.Update(id, blogUpdate)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully update blog",
		"blog":    blog.ToResponse(),
	})
}

func (bc *BlogController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	isDeleted := bc.Service.Repository.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "unsuccessfully delete blog",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully delete blog",
	})
}
