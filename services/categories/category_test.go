package categories

import (
	"mini-project/models"
	"mini-project/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoryRepository mocks.CategoryRepository
	categoryService    CategoryService

	categoryData    models.Category
	categoryRequest models.CategoryRequest
)

func TestMain(m *testing.M) {
	categoryService = NewCategoryService(&categoryRepository)

	categoryData = models.Category{
		ID:   1,
		Name: "category testing",
	}

	categoryRequest = models.CategoryRequest{
		Name: "category testing",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]models.Category{categoryData}).Once()

		result := categoryService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]models.Category{}).Once()

		result := categoryService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		categoryRepository.On("GetByID", 1).Return(categoryData).Once()

		result := categoryService.GetByID(1)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		categoryRepository.On("GetByID", -1).Return(models.Category{}).Once()

		result := categoryService.GetByID(-1)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		categoryRepository.On("Create", categoryRequest).Return(categoryData).Once()

		result := categoryService.Create(categoryRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		categoryRepository.On("Create", categoryRequest).Return(models.Category{}).Once()

		result := categoryService.Create(categoryRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		categoryRepository.On("Update", 1, categoryRequest).Return(categoryData).Once()

		result := categoryService.Update(1, categoryRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		categoryRepository.On("Update", -1, categoryRequest).Return(models.Category{}).Once()

		result := categoryService.Update(-1, categoryRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		categoryRepository.On("Delete", 1).Return(true).Once()

		result := categoryService.Delete(1)

		assert.Equal(t, true, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		categoryRepository.On("Delete", -1).Return(false).Once()

		result := categoryService.Delete(-1)

		assert.NotEqual(t, true, result)
	})
}
