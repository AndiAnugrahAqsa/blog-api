package services

import (
	"mini-project/models"
	"mini-project/repositories"
)

type CommentService struct {
	Repository repositories.CommentRepository
}

func NewCommentService() CommentService {
	return CommentService{
		Repository: &repositories.CommentRepositoryImpl{},
	}
}

func (cs *CommentService) GetAll() []models.Comment {
	return cs.Repository.GetAll()
}

func (cs *CommentService) GetByBlogID(blog_id int) []models.Comment {
	return cs.Repository.GetByBlogID(blog_id)
}

func (cs *CommentService) GetByID(id int) models.Comment {
	return cs.Repository.GetByID(id)
}

func (cs *CommentService) Create(commentRequest models.CommentRequest) models.Comment {
	return cs.Repository.Create(commentRequest)
}

func (cs *CommentService) Update(id int, commentRequest models.CommentRequest) models.Comment {
	return cs.Repository.Update(id, commentRequest)
}

func (cs *CommentService) Delete(id int) bool {
	return cs.Repository.Delete(id)
}
