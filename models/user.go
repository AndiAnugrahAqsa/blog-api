package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id"`
	IsAdmin   bool           `json:"is_admin"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Blogs     []Blog         `json:"blogs" gorm:"constraint:OnDelete:CASCADE;"`
	Likes     []Like         `json:"likes" gorm:"constraint:OnDelete:CASCADE;"`
	Comments  []Comment      `json:"comments" gorm:"constraint:OnDelete:CASCADE;"`
}

type UserRequest struct {
	IsAdmin   bool   `json:"is_admin"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (ur *UserRequest) ToDBForm() User {
	return User{
		IsAdmin:   ur.IsAdmin,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Email:     ur.Email,
		Password:  ur.Password,
	}
}

func (ur *UserRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(ur)

	return err
}

type UserResponse struct {
	ID        int            `json:"id"`
	IsAdmin   bool           `json:"is_admin"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		IsAdmin:   u.IsAdmin,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
