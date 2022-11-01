package services

import (
	"mini-project/models"
	"mini-project/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	roleRepository mocks.RoleRepository
	roleService    RoleService

	roleData    models.Role
	roleRequest models.RoleRequest
)

func TestMain(m *testing.M) {
	roleService = NewRoleService(&roleRepository)

	roleData = models.Role{
		ID:   1,
		Name: "role testing",
	}

	roleRequest = models.RoleRequest{
		Name: "role testing",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		roleRepository.On("GetAll").Return([]models.Role{roleData}).Once()

		result := roleService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		roleRepository.On("GetAll").Return([]models.Role{}).Once()

		result := roleService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		roleRepository.On("GetByID", 1).Return(roleData).Once()

		result := roleService.GetByID(1)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		roleRepository.On("GetByID", -1).Return(models.Role{}).Once()

		result := roleService.GetByID(-1)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		roleRepository.On("Create", roleRequest).Return(roleData).Once()

		result := roleService.Create(roleRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		roleRepository.On("Create", roleRequest).Return(models.Role{}).Once()

		result := roleService.Create(roleRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		roleRepository.On("Update", 1, roleRequest).Return(roleData).Once()

		result := roleService.Update(1, roleRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		roleRepository.On("Update", -1, roleRequest).Return(models.Role{}).Once()

		result := roleService.Update(-1, roleRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		roleRepository.On("Delete", 1).Return(true).Once()

		result := roleService.Delete(1)

		assert.Equal(t, true, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		roleRepository.On("Delete", -1).Return(false).Once()

		result := roleService.Delete(-1)

		assert.NotEqual(t, true, result)
	})
}
