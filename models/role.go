package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Role struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type RoleRequest struct {
	Name string `json:"name" validate:"required"`
}

func (rr RoleRequest) ToDBForm() Role {
	return Role{
		Name: rr.Name,
	}
}

func (rr *RoleRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(rr)

	return err
}
