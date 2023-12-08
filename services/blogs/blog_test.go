package blogs

import (
	"mini-project/models"
	"mini-project/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	blogRepository mocks.BlogRepository
	blogService    BlogService

	blogData    models.Blog
	blogRequest models.BlogRequest
)

func TestMain(m *testing.M) {
	blogService = NewBlogService(&blogRepository)

	user := models.User{
		ID:        1,
		IsAdmin:   true,
		FirstName: "first name",
		LastName:  "last name",
	}

	category := models.Category{
		ID:   1,
		Name: "category testing",
	}

	blogData = models.Blog{
		ID:         1,
		Content:    "Content testing",
		Title:      "title testing",
		UserID:     1,
		User:       user,
		CategoryID: 1,
		Category:   category,
	}

	blogRequest = models.BlogRequest{
		UserID:  1,
		Content: "Content testing",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		blogRepository.On("GetAll").Return([]models.Blog{blogData}).Once()

		result := blogService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		blogRepository.On("GetAll").Return([]models.Blog{}).Once()

		result := blogService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByUserID(t *testing.T) {
	t.Run("GetByUserID | Valid", func(t *testing.T) {
		blogRepository.On("GetByUserID", 1).Return([]models.Blog{blogData}).Once()

		result := blogService.GetByUserID(1)

		assert.Equal(t, 1, len(result))
	})

	t.Run("GetByUserID | InValid", func(t *testing.T) {
		blogRepository.On("GetByUserID", 1).Return([]models.Blog{}).Once()

		result := blogService.GetByUserID(1)

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByCategoryID(t *testing.T) {
	t.Run("GetByCategoryID | Valid", func(t *testing.T) {
		blogRepository.On("GetByCategoryID", 1).Return([]models.Blog{blogData}).Once()

		result := blogService.GetByCategoryID(1)

		assert.Equal(t, 1, len(result))
	})

	t.Run("GetByCategoryID | InValid", func(t *testing.T) {
		blogRepository.On("GetByCategoryID", 1).Return([]models.Blog{}).Once()

		result := blogService.GetByCategoryID(1)

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		blogRepository.On("GetByID", 1).Return(blogData).Once()

		result := blogService.GetByID(1)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		blogRepository.On("GetByID", -1).Return(models.Blog{}).Once()

		result := blogService.GetByID(-1)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		blogRepository.On("Create", blogRequest).Return(blogData).Once()

		result := blogService.Create(blogRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		blogRepository.On("Create", blogRequest).Return(models.Blog{}).Once()

		result := blogService.Create(blogRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		blogRepository.On("Update", 1, blogRequest).Return(blogData).Once()

		result := blogService.Update(1, blogRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		blogRepository.On("Update", -1, blogRequest).Return(models.Blog{}).Once()

		result := blogService.Update(-1, blogRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		blogRepository.On("Delete", 1).Return(true).Once()

		result := blogService.Delete(1)

		assert.Equal(t, true, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		blogRepository.On("Delete", -1).Return(false).Once()

		result := blogService.Delete(-1)

		assert.NotEqual(t, true, result)
	})
}
