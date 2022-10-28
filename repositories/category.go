package repositories

import (
	"mini-project/database"
	"mini-project/models"
)

type CategoryRepositoryImpl struct{}

func (cr *CategoryRepositoryImpl) GetAll() []models.Category {
	var categories []models.Category

	database.DB.Find(&categories)

	return categories
}

func (cr *CategoryRepositoryImpl) GetByID(id int) models.Category {
	var category models.Category

	database.DB.First(&category, id)

	return category
}

func (cr *CategoryRepositoryImpl) Create(categoryRequest models.CategoryRequest) models.Category {
	category := categoryRequest.ToDBForm()

	rec := database.DB.Create(&category)

	rec.Last(&category)

	return category
}

func (cr *CategoryRepositoryImpl) Update(id int, categoryRequest models.CategoryRequest) models.Category {
	category := cr.GetByID(id)

	category.Name = categoryRequest.Name

	rec := database.DB.Save(&category)

	rec.Last(&category)

	return category
}

func (cr *CategoryRepositoryImpl) Delete(id int) bool {
	category := cr.GetByID(id)

	rec := database.DB.Delete(&category)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
