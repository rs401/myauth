package models

import "gorm.io/gorm"

// User model for user
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password []byte `json:"-"`
}
