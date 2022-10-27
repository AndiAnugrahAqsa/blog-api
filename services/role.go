package services

import (
	"mini-project/models"
	"mini-project/repositories"
)

type RoleService struct {
	Repository repositories.RoleRepository
}

func NewRoleService() RoleService {
	return RoleService{
		Repository: &repositories.RoleRepositoryImpl{},
	}
}

func (cs *RoleService) GetAll() []models.Role {
	return cs.Repository.GetAll()
}

func (cs *RoleService) GetByID(id int) models.Role {
	return cs.Repository.GetByID(id)
}

func (cs *RoleService) Create(roleRequest models.RoleRequest) models.Role {
	return cs.Repository.Create(roleRequest)
}

func (cs *RoleService) Update(id int, roleRequest models.RoleRequest) models.Role {
	return cs.Repository.Update(id, roleRequest)
}

func (cs *RoleService) Delete(id int) bool {
	return cs.Repository.Delete(id)
}
