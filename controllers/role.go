package controllers

import (
	"mini-project/models"
	"mini-project/repositories"
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
		Service: services.NewRoleService(&repositories.RoleRepositoryImpl{}),
	}
}

func (rc *RoleController) GetAll(c echo.Context) error {
	var roles []models.Role
	roles = rc.Service.Repository.GetAll()

	return NewResponseSuccess(c, http.StatusOK, "successfully get all roles", roles)
}

func (rc *RoleController) GetByID(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var role models.Role

	role = rc.Service.Repository.GetByID(id)

	if role.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "role doesn't exist")
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get role", role)
}

func (rc *RoleController) Create(c echo.Context) error {
	var roleRequest models.RoleRequest

	c.Bind(&roleRequest)

	if err := roleRequest.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	role := rc.Service.Repository.Create(roleRequest)

	return NewResponseSuccess(c, http.StatusCreated, "successfully create role", role)
}

func (rc *RoleController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var roleUpdate models.RoleRequest

	c.Bind(&roleUpdate)

	if err := roleUpdate.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	role := rc.Service.Repository.Update(id, roleUpdate)

	if role.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "role doesn't exist")
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully update role", role)
}

func (rc *RoleController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	isDeleted := rc.Service.Repository.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "unsuccessfully delete role",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully delete role",
	})
}
