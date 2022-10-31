package repositories

import (
	"mini-project/database"
	"mini-project/models"

	"gorm.io/gorm/clause"
)

type RoleRepositoryImpl struct{}

func (rr *RoleRepositoryImpl) GetAll() []models.Role {
	var roles []models.Role

	database.DB.Preload(clause.Associations).Find(&roles)

	return roles
}

func (rr *RoleRepositoryImpl) GetByID(id int) models.Role {
	var role models.Role

	database.DB.Preload(clause.Associations).First(&role, id)

	return role
}

func (rr *RoleRepositoryImpl) Create(roleRequest models.RoleRequest) models.Role {
	role := roleRequest.ToDBForm()

	rec := database.DB.Create(&role)

	rec.Last(&role)

	return role
}

func (rr *RoleRepositoryImpl) Update(id int, roleRequest models.RoleRequest) models.Role {
	role := rr.GetByID(id)

	if role.ID == 0 {
		return role
	}

	role.Name = roleRequest.Name

	rec := database.DB.Save(&role)

	rec.Last(&role)

	return role
}

func (rr *RoleRepositoryImpl) Delete(id int) bool {
	role := rr.GetByID(id)

	rec := database.DB.Delete(&role)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
