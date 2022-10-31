package repositories

import (
	"mini-project/database"
	"mini-project/models"

	"gorm.io/gorm/clause"
)

type RoleRepositoryImpl struct{}

func (cr *RoleRepositoryImpl) GetAll() []models.Role {
	var roles []models.Role

	database.DB.Preload(clause.Associations).Find(&roles)

	return roles
}

func (cr *RoleRepositoryImpl) GetByID(id int) models.Role {
	var role models.Role

	database.DB.Preload(clause.Associations).First(&role, id)

	return role
}

func (cr *RoleRepositoryImpl) Create(roleRequest models.RoleRequest) models.Role {
	role := roleRequest.ToDBForm()

	rec := database.DB.Create(&role)

	rec.Last(&role)

	return role
}

func (cr *RoleRepositoryImpl) Update(id int, roleRequest models.RoleRequest) models.Role {
	role := cr.GetByID(id)

	if role.ID == 0 {
		return role
	}

	role.Name = roleRequest.Name

	rec := database.DB.Save(&role)

	rec.Last(&role)

	return role
}

func (cr *RoleRepositoryImpl) Delete(id int) bool {
	role := cr.GetByID(id)

	rec := database.DB.Delete(&role)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
