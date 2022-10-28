package services

import (
	"mini-project/models"
	"mini-project/repositories"
)

type CategoryService struct {
	Repository repositories.CategoryRepository
}

func NewCategoryService() CategoryService {
	return CategoryService{
		Repository: &repositories.CategoryRepositoryImpl{},
	}
}

func (cs *CategoryService) GetAll() []models.Category {
	return cs.Repository.GetAll()
}

func (cs *CategoryService) GetByID(id int) models.Category {
	return cs.Repository.GetByID(id)
}

func (cs *CategoryService) Create(categoryRequest models.CategoryRequest) models.Category {
	return cs.Repository.Create(categoryRequest)
}

func (cs *CategoryService) Update(id int, categoryRequest models.CategoryRequest) models.Category {
	return cs.Repository.Update(id, categoryRequest)
}

func (cs *CategoryService) Delete(id int) bool {
	return cs.Repository.Delete(id)
}
