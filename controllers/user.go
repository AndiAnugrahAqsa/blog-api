package controllers

import (
	"mini-project/middlewares"
	"mini-project/models"
	"mini-project/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service services.UserService
}

func NewUserController() UserController {
	return UserController{
		Service: services.NewUserService(),
	}
}

func (cc *UserController) GetAll(c echo.Context) error {
	var users []models.User
	users = cc.Service.Repository.GetAll()

	var usersResponse []models.UserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, user.ToResponse())
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get all users", usersResponse)
}

func (cc *UserController) GetByID(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var user models.User

	user = cc.Service.Repository.GetByID(id)

	return NewResponseSuccess(c, http.StatusOK, "successfully get user", user.ToResponse())
}

func (cc *UserController) Register(c echo.Context) error {
	var userRequest models.UserRequest

	c.Bind(&userRequest)

	user := cc.Service.Repository.Register(userRequest)

	return NewResponseSuccess(c, http.StatusCreated, "successfully register user", user.ToResponse())
}

func (cc *UserController) Login(c echo.Context) error {
	var userRequest models.UserRequest

	c.Bind(&userRequest)

	user := cc.Service.Repository.Login(userRequest)

	token := middlewares.GenerateToken(user)

	return c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (cc *UserController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var blogUpdate models.UserRequest

	c.Bind(&blogUpdate)

	user := cc.Service.Repository.Update(id, blogUpdate)

	return NewResponseSuccess(c, http.StatusOK, "successfully update user", user.ToResponse())
}

func (cc *UserController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	isDeleted := cc.Service.Repository.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "unsuccessfully delete user",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully delete user",
	})
}
