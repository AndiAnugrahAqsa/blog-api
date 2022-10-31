package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Category struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (cr CategoryRequest) ToDBForm() Category {
	return Category{
		Name: cr.Name,
	}
}

func (cr *CategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(cr)

	return err
}
