package users

import (
	"mini-project/models"
	"mini-project/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository mocks.UserRepository
	userService    UserService

	userData    models.User
	userRequest models.UserRequest
)

func TestMain(m *testing.M) {
	userService = NewUserService(&userRepository)

	roleData := models.Role{
		ID:   1,
		Name: "role testing",
	}

	userData = models.User{
		ID:        1,
		FirstName: "first name testing",
		LastName:  "last name testing",
		Email:     "testing@test.com",
		Password:  "123",
		RoleID:    1,
		Role:      roleData,
	}

	userRequest = models.UserRequest{
		FirstName: "first name testing",
		LastName:  "last name testing",
		RoleID:    1,
		Email:     "testing@test.com",
		Password:  "123",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		userRepository.On("GetAll").Return([]models.User{userData}).Once()

		result := userService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		userRepository.On("GetAll").Return([]models.User{}).Once()

		result := userService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		userRepository.On("GetByID", 1).Return(userData).Once()

		result := userService.GetByID(1)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		userRepository.On("GetByID", -1).Return(models.User{}).Once()

		result := userService.GetByID(-1)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		userRepository.On("Create", userRequest).Return(userData).Once()

		result := userService.Create(userRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		userRepository.On("Create", userRequest).Return(models.User{}).Once()

		result := userService.Create(userRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		userRepository.On("Login", userRequest).Return(userData).Once()

		result := userService.Login(userRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		userRepository.On("Login", userRequest).Return(models.User{}).Once()

		result := userService.Login(userRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		userRepository.On("Update", 1, userRequest).Return(userData).Once()

		result := userService.Update(1, userRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		userRepository.On("Update", -1, userRequest).Return(models.User{}).Once()

		result := userService.Update(-1, userRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		userRepository.On("Delete", 1).Return(true).Once()

		result := userService.Delete(1)

		assert.Equal(t, true, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		userRepository.On("Delete", -1).Return(false).Once()

		result := userService.Delete(-1)

		assert.NotEqual(t, true, result)
	})
}
