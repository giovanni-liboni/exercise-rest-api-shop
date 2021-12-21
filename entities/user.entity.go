package entities

import "time"

type User struct {
	ID        int64       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	Firstname string    `json:"firstname" db:"firstname"`
	Lastname  string    `json:"lastname" db:"lastname"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Role      string    `json:"role" db:"role"`
	Orders    []Order   `json:"orders"`
}

type UserForm struct {
	User
	PasswordForm string `json:"password"`
}

type UserStat struct {
	User
	TotalSpent float64   `json:"total_spent" db:"total_spent"`
}