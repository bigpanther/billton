package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name       string    `json:"name" gorm:"size:255;not null"`
	Username   string    `json:"username" gorm:"size:255;not null"`
	Email      string    `json:"email" gorm:"size:255;not null"`
	Role       string    `json:"role" gorm:"size:50;not null"`
	Warranties []Warranty
}

// IsConsumer checks if a user is a Consumer
func (u *User) IsConsumer() bool {
	return u.Role == UserRoleConsumer.String()
}

// IsAdmin checks if a user is a Admin
func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin.String()
}
