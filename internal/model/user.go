package model

import (
	"time"
)

// User is a user model for service layer.
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt *time.Time
}
