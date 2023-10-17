package model

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      int
	CreatedAt int64
	UpdatedAt int64
}
