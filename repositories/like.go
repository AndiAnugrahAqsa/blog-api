package repositories

import (
	"mini-project/database"
	"mini-project/models"

	"gorm.io/gorm/clause"
)

type LikeRepositoryImpl struct{}

func (lr *LikeRepositoryImpl) GetAll() []models.Like {
	var likes []models.Like

	database.DB.Preload(clause.Associations).Find(&likes)

	return likes
}

func (lr *LikeRepositoryImpl) GetByID(id int) models.Like {
	var like models.Like

	database.DB.Preload(clause.Associations).First(&like, id)

	return like
}

func (lr *LikeRepositoryImpl) GetByBlogID(blog_id int) []models.Like {
	var like []models.Like

	database.DB.Preload(clause.Associations).Find(&like, "blog_id = ?", blog_id)

	return like
}

func (lr *LikeRepositoryImpl) Create(likeRequest models.LikeRequest) models.Like {
	like := likeRequest.ToDBForm()

	rec := database.DB.Create(&like)

	rec.Last(&like)

	return like
}

func (lr *LikeRepositoryImpl) Delete(id int) bool {
	like := lr.GetByID(id)

	rec := database.DB.Delete(&like)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
