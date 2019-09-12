package models

import (
	"time"
)

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	ID        uint       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// User Model
type User struct {
	Model
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Address   string `gorm:"column:address" json:"address"`
	Email     string `gorm:"column:email" json:"email"`
}
