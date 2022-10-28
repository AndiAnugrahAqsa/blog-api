package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	User      User           `json:"user" gorm:"foreignKey:UserID;references:ID"`
	BlogID    int            `json:"blog_id"`
	Blog      Blog           `json:"blog" gorm:"foreignKey:BlogID;references:ID"`
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
	UserID  int    `json:"user_id"`
	BlogID  int    `json:"blog_id"`
	Content string `json:"content"`
}

func (cr CommentRequest) ToDBForm() Comment {
	return Comment{
		UserID:  cr.UserID,
		BlogID:  cr.BlogID,
		Content: cr.Content,
	}
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
