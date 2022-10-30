package repositories

import "mini-project/models"

type CategoryRepository interface {
	GetAll() []models.Category
	GetByID(id int) models.Category
	Create(categoryRequest models.CategoryRequest) models.Category
	Update(id int, categoryRequest models.CategoryRequest) models.Category
	Delete(id int) bool
}

type CommentRepository interface {
	GetAll() []models.Comment
	GetByBlogID(blog_id int) []models.Comment
	GetByID(id int) models.Comment
	Create(commentRequest models.CommentRequest) models.Comment
	Update(id int, commentRequest models.CommentRequest) models.Comment
	Delete(id int) bool
}

type LikeRepository interface {
	GetAll() []models.Like
	GetByBlogID(id int) []models.Like
	Create(likeRequest models.LikeRequest) models.Like
	Delete(id int) bool
}
type RoleRepository interface {
	GetAll() []models.Role
	GetByID(id int) models.Role
	Create(roleRequest models.RoleRequest) models.Role
	Update(id int, roleRequest models.RoleRequest) models.Role
	Delete(id int) bool
}

type UserRepository interface {
	GetAll() []models.User
	GetByID(id int) models.User
	Register(userRequest models.UserRequest) models.User
	Login(userRequest models.UserRequest) models.User
	Update(id int, userRequest models.UserRequest) models.User
	Delete(id int) bool
}
