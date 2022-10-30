package services

import (
	"mini-project/models"
	"mini-project/repositories"
)

type BlogService struct {
	Repository repositories.BlogRepository
}

func NewBlogService() BlogService {
	return BlogService{
		Repository: &repositories.BlogRepositoryImpl{},
	}
}

func (bs *BlogService) GetAll() []models.Blog {
	return bs.Repository.GetAll()
}

func (bs *BlogService) GetByUserID(user_id int) []models.Blog {
	return bs.Repository.GetByUserID(user_id)
}

func (bs *BlogService) GetByID(id int) models.Blog {
	return bs.Repository.GetByID(id)
}

func (bs *BlogService) Create(blogRequest models.BlogRequest) models.Blog {
	return bs.Repository.Create(blogRequest)
}

func (bs *BlogService) Update(id int, blogRequest models.BlogRequest) models.Blog {
	return bs.Repository.Update(id, blogRequest)
}

func (bs *BlogService) Delete(id int) bool {
	return bs.Repository.Delete(id)
}
