package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id" rw:"r"`
	Name      string    `json:"name" db:"name"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// IsConsumer checks if a user is a Consumer
func (u *User) IsConsumer() bool {
	return u.Role == UserRoleConsumer.String()
}

// IsAdmin checks if a user is a Admin
func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin.String()
}
