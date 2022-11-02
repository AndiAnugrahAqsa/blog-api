package repositories

import (
	"mini-project/database"
	"mini-project/models"

	"gorm.io/gorm/clause"
)

type BlogRepositoryImpl struct{}

var commentRepository CommentRepositoryImpl
var likeRepository LikeRepositoryImpl

func (br *BlogRepositoryImpl) GetAll() []models.Blog {
	var blogsFromDB []models.Blog
	var blogs []models.Blog

	database.DB.Preload(clause.Associations).Find(&blogsFromDB)

	for _, blog := range blogsFromDB {
		blog.Comments = commentRepository.GetByBlogID(blog.ID)
		blog.Likes = likeRepository.GetByBlogID(blog.ID)
		blogs = append(blogs, blog)
	}

	return blogs
}

func (br *BlogRepositoryImpl) GetByUserID(userID int) []models.Blog {
	var blogsFromDB []models.Blog
	var blogs []models.Blog

	database.DB.Preload(clause.Associations).Find(&blogsFromDB, "user_id = ?", userID)

	for _, blog := range blogsFromDB {
		blog.Comments = commentRepository.GetByBlogID(blog.ID)
		blog.Likes = likeRepository.GetByBlogID(blog.ID)
		blogs = append(blogs, blog)
	}

	return blogs
}

func (br *BlogRepositoryImpl) GetByCategoryID(categoryID int) []models.Blog {
	var blogsFromDB []models.Blog
	var blogs []models.Blog

	database.DB.Preload(clause.Associations).Find(&blogsFromDB, "category_id = ?", categoryID)

	for _, blog := range blogsFromDB {
		blog.Comments = commentRepository.GetByBlogID(blog.ID)
		blog.Likes = likeRepository.GetByBlogID(blog.ID)
		blogs = append(blogs, blog)
	}

	return blogs
}

func (br *BlogRepositoryImpl) GetByID(id int) models.Blog {
	var blog models.Blog

	database.DB.Preload(clause.Associations).First(&blog, id)

	blog.Comments = commentRepository.GetByBlogID(blog.ID)
	blog.Likes = likeRepository.GetByBlogID(blog.ID)

	return blog
}

func (br *BlogRepositoryImpl) Create(blogRequest models.BlogRequest) models.Blog {
	blog := blogRequest.ToDBForm()

	rec := database.DB.Create(&blog)

	rec.Last(&blog)

	return blog
}

func (br *BlogRepositoryImpl) Update(id int, blogRequest models.BlogRequest) models.Blog {
	blog := br.GetByID(id)

	if blog.ID == 0 {
		return blog
	}

	blog.UserID = blogRequest.UserID
	blog.CategoryID = blogRequest.CategoryID
	blog.Title = blogRequest.Title
	blog.Content = blogRequest.Content

	rec := database.DB.Save(&blog)

	rec.Last(&blog)

	return blog
}

func (br *BlogRepositoryImpl) Delete(id int) bool {
	blog := br.GetByID(id)

	rec := database.DB.Select("Likes", "Comments").Delete(&blog)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
