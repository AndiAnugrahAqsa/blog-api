package models

import (
	"time"

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
	Name string `json:"name"`
}

func (rr RoleRequest) ToDBForm() Role {
	return Role{
		Name: rr.Name,
	}
}
