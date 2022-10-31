package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Like struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	User      User           `json:"user"`
	BlogID    int            `json:"blog_id"`
	Blog      Blog           `json:"blog"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type LikeRequest struct {
	UserID int `json:"user_id" validate:"required"`
	BlogID int `json:"blog_id" validate:"required"`
}

func (lr *LikeRequest) ToDBForm() Like {
	return Like{
		UserID: lr.UserID,
		BlogID: lr.BlogID,
	}
}

func (lr *LikeRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(lr)

	return err
}

type LikeResponse struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	BlogID    int            `json:"blog_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (l *Like) ToResponse() LikeResponse {
	return LikeResponse{
		ID:        l.ID,
		UserID:    l.UserID,
		FirstName: l.User.FirstName,
		LastName:  l.User.LastName,
		BlogID:    l.BlogID,
		CreatedAt: l.CreatedAt,
		DeletedAt: l.DeletedAt,
	}
}
