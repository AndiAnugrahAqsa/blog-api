package comments

import (
	"mini-project/models"
	"mini-project/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	commentRepository mocks.CommentRepository
	commentService    CommentService

	commentData    models.Comment
	commentRequest models.CommentRequest
)

func TestMain(m *testing.M) {
	commentService = NewCommentService(&commentRepository)

	commentData = models.Comment{
		ID:      1,
		UserID:  1,
		BlogID:  1,
		Content: "Content testing",
	}

	commentRequest = models.CommentRequest{
		UserID:  1,
		BlogID:  1,
		Content: "Content testing",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		commentRepository.On("GetAll").Return([]models.Comment{commentData}).Once()

		result := commentService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		commentRepository.On("GetAll").Return([]models.Comment{}).Once()

		result := commentService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByBlogID(t *testing.T) {
	t.Run("GetByBlogID | Valid", func(t *testing.T) {
		commentRepository.On("GetByBlogID", 1).Return([]models.Comment{commentData}).Once()

		result := commentService.GetByBlogID(1)

		assert.Equal(t, 1, len(result))
	})

	t.Run("GetByBlogID | InValid", func(t *testing.T) {
		commentRepository.On("GetByBlogID", 1).Return([]models.Comment{}).Once()

		result := commentService.GetByBlogID(1)

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		commentRepository.On("GetByID", 1).Return(commentData).Once()

		result := commentService.GetByID(1)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		commentRepository.On("GetByID", -1).Return(models.Comment{}).Once()

		result := commentService.GetByID(-1)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		commentRepository.On("Create", commentRequest).Return(commentData).Once()

		result := commentService.Create(commentRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		commentRepository.On("Create", commentRequest).Return(models.Comment{}).Once()

		result := commentService.Create(commentRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		commentRepository.On("Update", 1, commentRequest).Return(commentData).Once()

		result := commentService.Update(1, commentRequest)

		assert.Equal(t, 1, result.ID)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		commentRepository.On("Update", -1, commentRequest).Return(models.Comment{}).Once()

		result := commentService.Update(-1, commentRequest)

		assert.NotEqual(t, 1, result.ID)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		commentRepository.On("Delete", 1).Return(true).Once()

		result := commentService.Delete(1)

		assert.Equal(t, true, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		commentRepository.On("Delete", -1).Return(false).Once()

		result := commentService.Delete(-1)

		assert.NotEqual(t, true, result)
	})
}
