package repositories

import "mini-project/models"

type RoleRepository interface {
	GetAll() []models.Role
	GetByID(id int) models.Role
	Create(roleRequest models.RoleRequest) models.Role
	Update(id int, roleRequest models.RoleRequest) models.Role
	Delete(id int) bool
}
