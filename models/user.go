package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id"`
	RoleID    int            `json:"role_id"`
	Role      Role           `json:"role"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserRequest struct {
	RoleID    int    `json:"role_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (ur *UserRequest) ToDBForm() User {
	return User{
		RoleID:    ur.RoleID,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Email:     ur.Email,
		Password:  ur.Password,
	}
}

type UserResponse struct {
	ID        int            `json:"id"`
	RoleName  string         `json:"role"`
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
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
