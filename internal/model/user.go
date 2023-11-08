package model

import (
	"time"
)

// User is a user model for service layer.
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type UserCredentials struct {
	Password        string
	ConfurmPassword string
}

type NewUser struct {
	User            *User
	UserCredentials *UserCredentials
}

// Validate validates the user model
func (u *NewUser) Validate(validator func(*NewUser) error) error {
	return validator(u)
}
