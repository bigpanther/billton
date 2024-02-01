package models

type User struct {
	Model
	// gorm.Model
	Name       string `json:"name" gorm:"size:255;not null"`
	Username   string `json:"username" gorm:"size:255;not null"`
	Email      string `json:"email" gorm:"size:255;not null"`
	Role       string `json:"role" gorm:"size:50;not null"`
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
