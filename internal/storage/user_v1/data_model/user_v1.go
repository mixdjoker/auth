package data_model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64 `db:"user_id"`
	Name      string
	Email     string
	Password  string
	Role      int
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
