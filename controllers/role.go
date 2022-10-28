package controllers

import (
	"mini-project/models"
	"mini-project/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleController struct {
	Service services.RoleService
}

func NewRoleController() RoleController {
	return RoleController{
		Service: services.NewRoleService(),
	}
}

func (cc *RoleController) GetAll(c echo.Context) error {
	var roles []models.Role
	roles = cc.Service.Repository.GetAll()

	return NewResponseSuccess(c, http.StatusOK, "successfully get all roles", roles)
}

func (cc *RoleController) GetByID(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var role models.Role

	role = cc.Service.Repository.GetByID(id)

	return NewResponseSuccess(c, http.StatusOK, "successfully get role", role)
}

func (cc *RoleController) Create(c echo.Context) error {
	var roleRequest models.RoleRequest

	c.Bind(&roleRequest)

	role := cc.Service.Repository.Create(roleRequest)

	return NewResponseSuccess(c, http.StatusCreated, "successfully create role", role)
}

func (cc *RoleController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var roleUpdate models.RoleRequest

	c.Bind(&roleUpdate)

	role := cc.Service.Repository.Update(id, roleUpdate)

	return NewResponseSuccess(c, http.StatusOK, "successfully update role", role)
}

func (cc *RoleController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	isDeleted := cc.Service.Repository.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "unsuccessfully delete role",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully delete role",
	})
}
