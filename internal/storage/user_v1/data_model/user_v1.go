package data_model

import (
	"database/sql"
	"time"
)

// User data model
type User struct {
	ID        int64 `db:"user_id"`
	Name      string
	Email     string
	Password  string
	Role      int          `db:"role_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
