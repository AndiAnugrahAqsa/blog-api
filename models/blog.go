package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID         int            `json:"id"`
	UserID     int            `json:"user_id"`
	User       User           `json:"user"`
	CategoryID int            `json:"category_id"`
	Category   Category       `json:"category"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	Comments   []Comment      `json:"comment"`
}

func (b *Blog) ToResponse() BlogResponse {
	var commentsResponse []CommentResponse

	for _, comment := range b.Comments {
		commentsResponse = append(commentsResponse, comment.ToResponse())
	}

	return BlogResponse{
		ID:           b.ID,
		UserID:       b.UserID,
		FirstName:    b.User.FirstName,
		LastName:     b.User.LastName,
		CategoryName: b.Category.Name,
		Title:        b.Title,
		Content:      b.Content,
		Comments:     commentsResponse,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}
}

type BlogRequest struct {
	UserID     int    `json:"user_id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

func (br *BlogRequest) ToDBForm() Blog {
	return Blog{
		UserID:     br.UserID,
		CategoryID: br.CategoryID,
		Title:      br.Title,
		Content:    br.Content,
	}
}

type BlogResponse struct {
	ID           int               `json:"id"`
	UserID       int               `json:"user_id"`
	FirstName    string            `json:"first_name"`
	LastName     string            `json:"last_name"`
	CategoryName string            `json:"category_name"`
	Title        string            `json:"title"`
	Content      string            `json:"content"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	DeletedAt    time.Time         `json:"deleted_at"`
	Comments     []CommentResponse `json:"comments"`
}
