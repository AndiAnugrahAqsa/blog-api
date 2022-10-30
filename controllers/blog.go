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

	return NewResponseSuccess(c, http.StatusOK, "successfully get all blogs", blogsResponse)
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

	return NewResponseSuccess(c, http.StatusOK, "successfully get blogs by user id", blogsResponse)
}

func (bc *BlogController) GetByID(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var blog models.Blog

	blog = bc.Service.Repository.GetByID(id)

	return NewResponseSuccess(c, http.StatusOK, "successfully get blog", blog.ToResponse())
}

func (bc *BlogController) Create(c echo.Context) error {
	var blogRequest models.BlogRequest

	c.Bind(&blogRequest)

	blog := bc.Service.Repository.Create(blogRequest)

	return NewResponseSuccess(c, http.StatusOK, "successfully create blog", blog.ToResponse())
}

func (bc *BlogController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var blogUpdate models.BlogRequest

	c.Bind(&blogUpdate)

	blog := bc.Service.Repository.Update(id, blogUpdate)

	return NewResponseSuccess(c, http.StatusOK, "successfully update blog", blog.ToResponse())
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
