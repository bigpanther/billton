package models

// UserRole represents the UserRole enum
type UserRole string

const (

	// UserRoleAdmin represents Admin UserRole
	UserRoleAdmin UserRole = "Admin"
	// UserRoleConsumer represents Consumer UserRole
	UserRoleConsumer UserRole = "Consumer"
)

// String returns the string representation of
func (k UserRole) String() string {
	return string(k)
}

// IsValidUserRole validates if the input is a UserRole
func IsValidUserRole(s string) bool {
	t := UserRole(s)
	return UserRoleAdmin == t || UserRoleConsumer == t
}
