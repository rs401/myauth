package models

import (
	"time"

	"gorm.io/gorm"
)

// User model for user
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  []byte `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
