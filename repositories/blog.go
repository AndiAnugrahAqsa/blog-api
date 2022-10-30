package repositories

import (
	"log"
	"mini-project/database"
	"mini-project/models"
)

type BlogRepositoryImpl struct{}

var commentRepository CommentRepositoryImpl

func (br *BlogRepositoryImpl) GetAll() []models.Blog {
	var blogsFromDB []models.Blog
	var blogs []models.Blog

	database.DB.Preload("User").Preload("Category").Find(&blogsFromDB)

	for _, blog := range blogsFromDB {
		blog.Comments = commentRepository.GetByBlogID(blog.ID)
		blogs = append(blogs, blog)
	}

	return blogs
}

func (br *BlogRepositoryImpl) GetAllWithComments() []models.Blog {
	var blogs []models.Blog
	// var blogsRes []models.BlogResponse
	// var comments []models.Comment
	// commentRepo := CommentRepositoryImpl{}
	// userRepo := UserRepositoryImpl{}
	// var comments = commentRepo.GetAll()
	// users := userRepo.GetAll()

	// database.DB.Preload(clause.Associations).Find(&blogs)
	// .Joins("JOIN credit_cards ON credit_cards.user_id = users.id")
	// database.DB.Preload(clause.Associations).Joins("LEFT JOIN comments ON comments.blog_id = blogs.id").Joins("LEFT JOIN SELECT * FROM comments LEFT JOIN comment ON comments.user_id = users.id").Preload("User").Find(&blogs)
	// database.DB.Preload("Comment").Joins("INNER JOIN users ON comments.user_id = users.id").Find(&blogs)

	// database.DB.Joins("Comment").Find(&blogs)
	// blogs = br.GetAll()
	// database.DB.Model(&blogs).Association("Comments").Find(&comments)
	var results struct {
		ID        int
		Content   string
		Comment   []models.Comment
		UserID    int
		FirstName string
		LastName  string
	}

	query := database.DB.Table("comments").Select("comments.id as ID, content, blog_id, users.first_name as first_name").Joins("left join users on comments.user_id = users.id")
	database.DB.Model(&blogs).Joins("join (?) q on blogs.id = q.blog_id", query).Scan(&results)

	log.Println(blogs)
	log.Println("res", results)

	return blogs
}

func (br *BlogRepositoryImpl) GetByID(id int) models.Blog {
	var blog models.Blog

	database.DB.Preload("User").Preload("Category").Preload("Comment").Preload("Comment.User").First(&blog, id)

	// blog.Comments = commentRepository.GetByBlogID(blog.ID)

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

	rec := database.DB.Delete(&blog)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
