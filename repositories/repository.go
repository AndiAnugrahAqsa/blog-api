package repositories

import "mini-project/models"

type RoleRepository interface {
	GetAll() []models.Role
	GetByID(id int) models.Role
	Create(roleRequest models.RoleRequest) models.Role
	Update(id int, roleRequest models.RoleRequest) models.Role
	Delete(id int) bool
}

type UserRepository interface {
	GetAll() []models.User
	GetByID(id int) models.User
	Register(userRequest models.UserRequest) models.User
	Login(userRequest models.UserRequest) models.User
	Update(id int, userRequest models.UserRequest) models.User
	Delete(id int) bool
}
