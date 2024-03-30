package domain

import (
	"time"
)

type UserRoles string

const (
	AdminRole UserRoles = "admin"
	UserRole  UserRoles = "user"
)

type User struct {
	Id        int
	FullName  string
	Email     string
	Phone     string
	Password  string
	Role      UserRoles
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
