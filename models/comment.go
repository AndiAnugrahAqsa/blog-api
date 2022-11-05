package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Comment struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	User      User           `json:"user"`
	BlogID    int            `json:"blog_id"`
	Blog      Blog           `json:"blog"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (c Comment) ToResponse() CommentResponse {
	return CommentResponse{
		ID:        c.ID,
		UserID:    c.UserID,
		FirstName: c.User.FirstName,
		LastName:  c.User.LastName,
		BlogID:    c.BlogID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}

type CommentRequest struct {
	UserID  int    `json:"user_id" validate:"required"`
	BlogID  int    `json:"blog_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (cr CommentRequest) ToDBForm() Comment {
	return Comment{
		UserID:  cr.UserID,
		BlogID:  cr.BlogID,
		Content: cr.Content,
	}
}

func (cr *CommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(cr)

	return err
}

type CommentResponse struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	BlogID    int            `json:"blog_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
