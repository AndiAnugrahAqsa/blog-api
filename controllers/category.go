package controllers

import (
	"mini-project/models"
	"mini-project/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	Service services.CategoryService
}

func NewCategoryController() CategoryController {
	return CategoryController{
		Service: services.NewCategoryService(),
	}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	var categories []models.Category

	categories = cc.Service.Repository.GetAll()

	return NewResponseSuccess(c, http.StatusOK, "successfully get all categories", categories)
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var category models.Category

	category = cc.Service.Repository.GetByID(id)

	return NewResponseSuccess(c, http.StatusOK, "successfully get category", category)
}

func (cc *CategoryController) Create(c echo.Context) error {
	var categoryRequest models.CategoryRequest

	c.Bind(&categoryRequest)

	if err := categoryRequest.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	category := cc.Service.Repository.Create(categoryRequest)

	return NewResponseSuccess(c, http.StatusCreated, "successfully create category", category)
}

func (cc *CategoryController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var categoryUpdate models.CategoryRequest

	c.Bind(&categoryUpdate)

	if err := categoryUpdate.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	category := cc.Service.Repository.Update(id, categoryUpdate)

	return NewResponseSuccess(c, http.StatusOK, "successfully update category", category)
}

func (cc *CategoryController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	isDeleted := cc.Service.Repository.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "unsuccessfully delete category",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully delete category",
	})
}
