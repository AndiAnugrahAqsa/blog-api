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

func (bs *BlogService) GetByUserID(userID int) []models.Blog {
	return bs.Repository.GetByUserID(userID)
}

func (bs *BlogService) GetByCategoryID(categoryID int) []models.Blog {
	return bs.Repository.GetByCategoryID(categoryID)
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
