package entities

import "time"

type User struct {
	ID        int     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-" db:"password"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `json:"role"`
}
