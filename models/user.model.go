package models

import (
	"time"
)

// UserModel represents a system user.
// @Description Model that includes basic user information such as name, surname, username, email and password.
type UserModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null"`
	Surname   string    `gorm:"not null"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// UserCreateModel para la creación de usuarios
type UserCreateModel struct {
	Name     string `json:"name" example:"John"`
	Surname  string `json:"surname" example:"Doe"`
	Username string `json:"username" example:"john_doe"`
	Email    string `json:"email" example:"johndoe@example.com"`
	Password string `json:"password" example:"verySecureP4ssword"`
}

// UserUpdateModel para actualizar usuarios
type UserUpdateModel struct {
	Name     string `json:"name,omitempty" example:"John"`
	Surname  string `json:"surname,omitempty" example:"Doe"`
	Username string `json:"username,omitempty" example:"john_doe"`
	Email    string `json:"email,omitempty" example:"johndoe@example.com"`
}

// UserViewModel para visualizar usuarios
type UserViewModel struct {
	ID        uint      `json:"id" example:"1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" example:"John"`
	Surname   string    `json:"surname" example:"Doe"`
	Username  string    `json:"username" example:"john_doe"`
	Email     string    `json:"email" example:"johndoe@example.com"`
}
