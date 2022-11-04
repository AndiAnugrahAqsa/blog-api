package repositories

import (
	"mini-project/database"
	"mini-project/models"

	"gorm.io/gorm/clause"
)

type CommentRepositoryImpl struct{}

func (cr *CommentRepositoryImpl) GetAll() []models.Comment {
	var comments []models.Comment

	database.DB.Preload(clause.Associations).Find(&comments)

	return comments
}

func (cr *CommentRepositoryImpl) GetByBlogID(blog_id int) []models.Comment {
	var comments []models.Comment

	database.DB.Preload(clause.Associations).Find(&comments, "blog_id = ?", blog_id)

	return comments
}

func (cr *CommentRepositoryImpl) GetByID(id int) models.Comment {
	var comment models.Comment

	database.DB.Preload(clause.Associations).First(&comment, id)

	return comment
}

func (cr *CommentRepositoryImpl) Create(commentRequest models.CommentRequest) models.Comment {
	comment := commentRequest.ToDBForm()

	rec := database.DB.Create(&comment)

	rec.Preload(clause.Associations).Last(&comment)

	return comment
}

func (cr *CommentRepositoryImpl) Update(id int, commentRequest models.CommentRequest) models.Comment {
	comment := cr.GetByID(id)

	if comment.ID == 0 {
		return comment
	}

	comment.UserID = commentRequest.UserID
	comment.BlogID = commentRequest.BlogID
	comment.Content = commentRequest.Content

	rec := database.DB.Save(&comment)

	rec.Preload(clause.Associations).Last(&comment)

	return comment
}

func (cr *CommentRepositoryImpl) Delete(id int) bool {
	comment := cr.GetByID(id)

	rec := database.DB.Delete(&comment)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
