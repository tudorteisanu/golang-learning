package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name     string     `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
