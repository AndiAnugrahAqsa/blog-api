package likes

import (
	"mini-project/models"
	"mini-project/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	likeRepository mocks.LikeRepository
	likeService    LikeService

	likeData    models.Like
	likeRequest models.LikeRequest
)

func TestMain(m *testing.M) {
	likeService = NewLikeService(&likeRepository)

	blog := models.Blog{
		ID:         1,
		UserID:     1,
		CategoryID: 1,
		Title:      "title testing",
		Content:    "content testing",
	}

	user := models.User{
		ID:        1,
		IsAdmin:   true,
		FirstName: "first name",
		LastName:  "last name",
	}

	likeData = models.Like{
		ID:     1,
		BlogID: 1,
		Blog:   blog,
		User:   user,
	}

	likeRequest = models.LikeRequest{
		UserID: 1,
		BlogID: 1,
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		likeRepository.On("GetAll").Return([]models.Like{likeData}).Once()

		result := likeService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		likeRepository.On("GetAll").Return([]models.Like{}).Once()

		result := likeService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByBlogID(t *testing.T) {
	t.Run("Get ByBlogID | Valid", func(t *testing.T) {
		likeRepository.On("GetByBlogID", 1).Return([]models.Like{likeData}).Once()

		result := likeService.GetByBlogID(1)

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get ByBlogID | InValid", func(t *testing.T) {
		likeRepository.On("GetByBlogID", 1).Return([]models.Like{}).Once()

		result := likeService.GetByBlogID(1)

		assert.Equal(t, 0, len(result))
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		likeRepository.On("Create", likeRequest).Return(likeData).Once()

		result := likeService.Create(likeRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		likeRepository.On("Create", likeRequest).Return(models.Like{}).Once()

		result := likeService.Create(likeRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		likeRepository.On("Delete", 1).Return(true).Once()

		result := likeService.Delete(1)

		assert.Equal(t, true, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		likeRepository.On("Delete", -1).Return(false).Once()

		result := likeService.Delete(-1)

		assert.NotEqual(t, true, result)
	})
}
