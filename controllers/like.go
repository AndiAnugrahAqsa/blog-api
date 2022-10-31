package controllers

import (
	"mini-project/models"
	"mini-project/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LikeController struct {
	Service services.LikeService
}

func NewLikeController() LikeController {
	return LikeController{
		Service: services.NewLikeService(),
	}
}

func (lc *LikeController) GetAll(c echo.Context) error {
	var likes []models.Like
	likes = lc.Service.Repository.GetAll()

	var likesResponse []models.LikeResponse

	for _, like := range likes {
		likesResponse = append(likesResponse, like.ToResponse())
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get all likes", likesResponse)
}

func (lc *LikeController) GetByBlogID(c echo.Context) error {
	blogIDString := c.Param("blog_id")
	blogID, _ := strconv.Atoi(blogIDString)

	var likes []models.Like

	likes = lc.Service.Repository.GetByBlogID(blogID)

	var likesResponse []models.LikeResponse

	for _, like := range likes {
		likesResponse = append(likesResponse, like.ToResponse())
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get likes by blog id", likesResponse)
}

func (lc *LikeController) Create(c echo.Context) error {
	var likeRequest models.LikeRequest

	c.Bind(&likeRequest)

	if err := likeRequest.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	like := lc.Service.Repository.Create(likeRequest)

	return NewResponseSuccess(c, http.StatusOK, "successfully create like", like.ToResponse())
}

func (lc *LikeController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	isDeleted := lc.Service.Repository.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "unsuccessfully delete like",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully delete like",
	})
}
