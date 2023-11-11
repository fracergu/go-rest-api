package models

import (
	"time"

	"gorm.io/gorm"
)

// UserModel represents a system user.
// @Description Model that includes basic user information such as name, surname, username, email and password.
type UserModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name" gorm:"not null"`
	Surname   string         `json:"surname" gorm:"not null"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
}
