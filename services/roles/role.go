package roles

import (
	"mini-project/models"
	"mini-project/repositories"
)

type RoleService struct {
	Repository repositories.RoleRepository
}

func NewRoleService(repository repositories.RoleRepository) RoleService {
	return RoleService{
		Repository: repository,
	}
}

func (rs *RoleService) GetAll() []models.Role {
	return rs.Repository.GetAll()
}

func (rs *RoleService) GetByID(id int) models.Role {
	return rs.Repository.GetByID(id)
}

func (rs *RoleService) Create(roleRequest models.RoleRequest) models.Role {
	return rs.Repository.Create(roleRequest)
}

func (rs *RoleService) Update(id int, roleRequest models.RoleRequest) models.Role {
	return rs.Repository.Update(id, roleRequest)
}

func (rs *RoleService) Delete(id int) bool {
	return rs.Repository.Delete(id)
}
